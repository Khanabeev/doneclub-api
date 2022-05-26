package goal

import (
	"doneclub-api/internal/adapters/api"
	"doneclub-api/internal/domain/goal"
	"doneclub-api/pkg/apperrors"
	"encoding/json"
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
	service goal.Service
}

func NewHandler(service goal.Service) api.Handler {
	return &handler{service: service}
}

func (h *handler) Register(router *mux.Router) {
	router.
		HandleFunc(createGoalUrl, h.CreateGoal).
		Methods(http.MethodPost).
		Name("CreateGoal")
}

func (h handler) CreateGoal(w http.ResponseWriter, r *http.Request) {
	var dto goal.RequestCreateGoalDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		api.WriteResponse(w, http.StatusBadRequest, apperrors.NewBadRequest("Incorrect request"))
	} else {
		response, appErr := h.service.CreateNewGoal(r.Context(), &dto)
		if appErr != nil {
			api.WriteResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			api.WriteResponse(w, http.StatusOK, response)
		}
	}
}
