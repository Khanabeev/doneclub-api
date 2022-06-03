package goal

import (
	"doneclub-api/internal/adapters/api"
	"doneclub-api/internal/domain/goal"
	"doneclub-api/pkg/apperrors"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const (
	createGoalUrl  = "/api/goals"
	getAllGoalsUrl = "/api/goals/all"
	getGoalUrl     = "/api/goals/{goal_id:[0-9]+}"
	updateGoalUrl  = "/api/goals/{goal_id:[0-9]+}"
	deleteGoalUrl  = "/api/goals/{goal_id:[0-9]+}"
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
		HandleFunc(getGoalUrl, h.GetGoal).
		Methods(http.MethodGet).
		Name("GetGoal")

	router.
		HandleFunc(getAllGoalsUrl, h.GetAllGoals).
		Methods(http.MethodGet).
		Name("GetAllGoals")

	router.
		HandleFunc(deleteGoalUrl, h.DeleteGoal).
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
		goalId, err := strconv.Atoi(vars["goal_id"])
		if err != nil {
			newErr := apperrors.NewUnexpectedError("Unexpected error")
			api.WriteResponse(w, newErr.Code, newErr.AsMessage())
		}
		response, appErr := h.service.UpdateGoal(r.Context(), &dto, goalId)
		if appErr != nil {
			api.WriteResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			api.WriteResponse(w, http.StatusOK, response)
		}
	}
}

func (h handler) GetGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	goalId, err := strconv.Atoi(vars["goal_id"])
	if err != nil {
		newErr := apperrors.NewUnexpectedError("Unexpected error")
		api.WriteResponse(w, newErr.Code, newErr.AsMessage())
	}

	response, appErr := h.service.GetGoal(r.Context(), goalId)
	if appErr != nil {
		api.WriteResponse(w, appErr.Code, appErr.AsMessage())
	} else {
		api.WriteResponse(w, http.StatusOK, response)
	}
}

func (h handler) GetAllGoals(w http.ResponseWriter, r *http.Request) {
	status, _ := r.URL.Query()["status"]

	response, appErr := h.service.GetAllGoals(r.Context(), status[0])
	if appErr != nil {
		api.WriteResponse(w, appErr.Code, appErr.AsMessage())
	} else {
		api.WriteResponse(w, http.StatusOK, response)
	}
}

func (h handler) DeleteGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	goalId, err := strconv.Atoi(vars["goal_id"])
	if err != nil {
		newErr := apperrors.NewUnexpectedError("Unexpected error")
		api.WriteResponse(w, newErr.Code, newErr.AsMessage())
	}

	response, appErr := h.service.DeleteGoal(r.Context(), goalId)
	if appErr != nil {
		api.WriteResponse(w, appErr.Code, appErr.AsMessage())
	} else {
		api.WriteResponse(w, http.StatusOK, response)
	}
}
