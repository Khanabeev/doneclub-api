package user

import (
	"doneclub-api/internal/adapters/api"
	"doneclub-api/internal/domain/user"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	userProfileUrl = "/api/users/profile"
)

type handler struct {
	service user.Service
}

func (h *handler) Register(router *mux.Router) {
	router.
		HandleFunc(userProfileUrl, h.GetUserProfile).
		Methods(http.MethodGet).
		Name("GetUserProfile")
}

func NewHandler(service user.Service) api.Handler {
	return &handler{service: service}
}

func (h handler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	response, err := h.service.FindUserById(r.Context())
	if err != nil {
		api.WriteResponse(w, err.Code, err.AsMessage())
	} else {
		api.WriteResponse(w, http.StatusOK, response)
	}
}
