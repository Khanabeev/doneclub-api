package user

import (
	"context"
	"doneclub-api/internal/adapters/api"
	"doneclub-api/internal/domain/user"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	usersUrl = "/api/users"
	//userUrl  = "/api/users/:user_id"
)

type handler struct {
	service user.Service
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, usersUrl, h.CreateUser)
}

func NewHandler(service user.Service) api.Handler {
	return &handler{service: service}
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var request user.CreateUserRequestDTO
	ctx := context.Background()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		api.WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		response, err := h.service.CreateUser(ctx, &request)
		if err != nil {
			api.WriteResponse(w, err.Code, err.AsMessage())
		} else {
			api.WriteResponse(w, http.StatusOK, response)
		}
	}
}
