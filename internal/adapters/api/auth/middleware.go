package auth

import (
	"context"
	"doneclub-api/internal/adapters/api"
	"doneclub-api/internal/domain/auth"
	"doneclub-api/pkg/apperrors"
	"doneclub-api/pkg/logging"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"time"
)

type Middleware struct {
	Storage auth.Storage
}

type UserClaims struct {
	ID    int
	Email string
	Role  string
}

type ContextKey string

const ContextUserKey ContextKey = "user"

func (a Middleware) AuthorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			currentRouteVars := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")

			if authHeader != "" {
				token := getTokenFromHeader(authHeader)

				isAuthorized := a.Storage.IsAuthorized(token, currentRoute.GetName(), currentRouteVars)

				if isAuthorized {
					claims, err := ClaimsFromToken(token)
					if err != nil {
						logger := logging.GetLogger()
						logger.Error(err)
						appError := apperrors.NewUnauthorizedError("Incorrect token claims").AsMessage()
						api.WriteResponse(w, appError.Code, appError.AsMessage())
					}
					ctx, cancel := context.WithTimeout(r.Context(), time.Duration(60*time.Second))
					defer cancel()

					u := &UserClaims{
						ID:    claims.UserID,
						Email: claims.Email,
						Role:  claims.Role,
					}

					ctx = context.WithValue(r.Context(), ContextUserKey, u)
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					appError := apperrors.NewUnauthorizedError("Unauthorized")
					api.WriteResponse(w, appError.Code, appError.AsMessage())
				}
			} else {
				appError := apperrors.NewUnauthorizedError("missing token").AsMessage()
				api.WriteResponse(w, appError.Code, appError.AsMessage())
			}
		})
	}
}

func getTokenFromHeader(header string) string {
	/*
	   token is coming in the format as below
	   "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50cyI6W.yI5NTQ3MCIsIjk1NDcyIiw"
	*/
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
