package task

import (
	"doneclub-api/internal/adapters/api"
	"doneclub-api/internal/domain/task"
	"doneclub-api/pkg/apperrors"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const (
	createTaskUrl     = "/api/tasks"
	getAllTasksUrl    = "/api/tasks"
	getAllTasksByGoal = "/api/goals/{goal_id:[0-9]+}/tasks"
	getTaskUrl        = "/api/tasks/{task_id:[0-9]+}"
	updateTaskUrl     = "/api/tasks/{task_id:[0-9]+}"
	deleteTaskUrl     = "/api/tasks/{task_id:[0-9]+}"
)

type handler struct {
	service Service
}

func NewHandler(service Service) api.Handler {
	return &handler{service: service}
}

func (h *handler) Register(router *mux.Router) {
	router.
		HandleFunc(createTaskUrl, h.CreateTask).
		Methods(http.MethodPost).
		Name("CreateTask")

	router.
		HandleFunc(getAllTasksUrl, h.GetAllTasks).
		Methods(http.MethodGet).
		Name("GetAllTasks")

	router.
		HandleFunc(getAllTasksUrl, h.GetAllTasksByGoal).
		Methods(http.MethodGet).
		Name("GetAllTasksByGoal")
}

func (h *handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var dto task.RequestCreateTaskDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		fmt.Println(err)
		api.WriteResponse(w, http.StatusBadRequest, apperrors.NewBadRequest("Incorrect request"))
	} else {
		response, appErr := h.service.CreateNewTask(r.Context(), &dto)
		if appErr != nil {
			api.WriteResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			api.WriteResponse(w, http.StatusOK, response)
		}
	}
}

func (h *handler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	status := api.GetStatus(r)
	response, appErr := h.service.GetAllTasks(r.Context(), status)
	if appErr != nil {
		api.WriteResponse(w, appErr.Code, appErr.AsMessage())
	} else {
		api.WriteResponse(w, http.StatusOK, response)
	}
}

func (h *handler) GetAllTasksByGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	goalId, err := strconv.Atoi(vars["goal_id"])
	if err != nil {
		newErr := apperrors.NewUnexpectedError("Unexpected error")
		api.WriteResponse(w, newErr.Code, newErr.AsMessage())
	}

	status := api.GetStatus(r)
	response, appErr := h.service.GetAllTasksByGoal(r.Context(), status, goalId)
	if appErr != nil {
		api.WriteResponse(w, appErr.Code, appErr.AsMessage())
	} else {
		api.WriteResponse(w, http.StatusOK, response)
	}
}
