package user

import (
	"doneclub-api/internal/adapters/api"
	"doneclub-api/internal/domain/user"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	usersUrl = "/api/users"
	userUrl  = "/api/users/{user_id:[0-9]+}"
)

type handler struct {
	service user.Service
}

func (h *handler) Register(router *mux.Router) {
	router.
		HandleFunc(userUrl, h.GetUser).
		Methods(http.MethodGet).
		Name("GetUser")

	router.
		HandleFunc(usersUrl, h.GetUsers).
		Methods(http.MethodGet).
		Name("GetUsers")
}

func NewHandler(service user.Service) api.Handler {
	return &handler{service: service}
}

func (h handler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	api.WriteResponse(w, http.StatusOK, vars)
}

func (h handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	vars := []string{"Hello world"}
	api.WriteResponse(w, http.StatusOK, vars)
}
