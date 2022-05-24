package goal

import (
	"doneclub-api/internal/adapters/api"
	"doneclub-api/internal/domain/user"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	createGoalUrl = "/api/goals"
	getGoalUrl    = "/api/goals/{goal_id:[0-9]+}"
	updateGoalUrl = "/api/goals/{goal_id:[0-9]+}"
	deleteGoalUrl = "/api/goals/{goal_id:[0-9]+}"
)

type handler struct {
	service user.Service
}

func NewHandler(service user.Service) api.Handler {
	return &handler{service: service}
}

func (h *handler) Register(router *mux.Router) {
	router.
		HandleFunc(createGoalUrl, h.CreateGoal).
		Methods(http.MethodPost).
		Name("CreateGoal")
}

func (h handler) CreateGoal(w http.ResponseWriter, r *http.Request) {
	api.WriteResponse(w, http.StatusOK, "hello")
}
