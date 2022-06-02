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
	service Service
}

func NewHandler(service Service) api.Handler {
	return &handler{service: service}
}

func (h *handler) Register(router *mux.Router) {
	router.
		HandleFunc(createGoalUrl, h.CreateGoal).
		Methods(http.MethodPost).
		Name("CreateGoal")

	router.
		HandleFunc(updateGoalUrl, h.UpdateGoal).
		Methods(http.MethodPut).
		Name("UpdateGoal")

	router.
		HandleFunc(getGoalUrl, h.CreateGoal).
		Methods(http.MethodGet).
		Name("GetGoal")

	router.
		HandleFunc(deleteGoalUrl, h.CreateGoal).
		Methods(http.MethodDelete).
		Name("DeleteGoal")
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

func (h handler) UpdateGoal(w http.ResponseWriter, r *http.Request) {
	var dto goal.RequestUpdateGoalDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		api.WriteResponse(w, http.StatusBadRequest, apperrors.NewBadRequest("Incorrect request"))
	} else {
		vars := mux.Vars(r)
		goalId := vars["goal_id"]
		response, appErr := h.service.UpdateGoal(r.Context(), &dto, goalId)
		if appErr != nil {
			api.WriteResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			api.WriteResponse(w, http.StatusOK, response)
		}
	}
}
