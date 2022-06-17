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
	updateTaskGoalUrl = "/api/tasks/{task_id:[0-9]+}/goal/{goal_id:[0-9]+}"
	deleteTaskGoalUrl = "/api/tasks/{task_id:[0-9]+}/goal"
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
		HandleFunc(getAllTasksByGoal, h.GetAllTasksByGoal).
		Methods(http.MethodGet).
		Name("GetAllTasksByGoal")

	router.
		HandleFunc(getTaskUrl, h.GetTaskById).
		Methods(http.MethodGet).
		Name("GetTaskById")

	router.
		HandleFunc(updateTaskUrl, h.UpdateTask).
		Methods(http.MethodPut).
		Name("UpdateTask")

	router.
		HandleFunc(updateTaskGoalUrl, h.UpdateTaskGoal).
		Methods(http.MethodPut).
		Name("UpdateTaskGoal")

	router.
		HandleFunc(deleteTaskGoalUrl, h.DeleteTaskGoal).
		Methods(http.MethodDelete).
		Name("DeleteTaskGoal")

	router.
		HandleFunc(deleteTaskUrl, h.DeleteTask).
		Methods(http.MethodDelete).
		Name("DeleteTask")
}

func (h *handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var dto task.RequestCreateTaskDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
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

func (h *handler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["task_id"])
	if err != nil {
		newErr := apperrors.NewUnexpectedError("Unexpected error")
		api.WriteResponse(w, newErr.Code, newErr.AsMessage())
	}

	response, appErr := h.service.GetTaskById(r.Context(), taskId)
	if appErr != nil {
		api.WriteResponse(w, appErr.Code, appErr.AsMessage())
	} else {
		api.WriteResponse(w, http.StatusOK, response)
	}
}

func (h *handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var dto task.RequestUpdateTaskDTO
	d := json.NewDecoder(r.Body)
	d.UseNumber()
	err := d.Decode(&dto)
	if err != nil {
		api.WriteResponse(w, http.StatusBadRequest, apperrors.NewBadRequest(fmt.Sprintf("Incorrect request: %s", err.Error())))
	} else {
		vars := mux.Vars(r)
		taskId, err := strconv.Atoi(vars["task_id"])
		if err != nil {
			newErr := apperrors.NewUnexpectedError("Unexpected error")
			api.WriteResponse(w, newErr.Code, newErr.AsMessage())
		}
		response, appErr := h.service.UpdateTask(r.Context(), &dto, taskId)
		if appErr != nil {
			api.WriteResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			api.WriteResponse(w, http.StatusOK, response)
		}
	}
}

func (h *handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["task_id"])
	if err != nil {
		newErr := apperrors.NewUnexpectedError("Unexpected error")
		api.WriteResponse(w, newErr.Code, newErr.AsMessage())
	}

	response, appErr := h.service.DeleteTask(r.Context(), taskId)
	if appErr != nil {
		api.WriteResponse(w, appErr.Code, appErr.AsMessage())
	} else {
		api.WriteResponse(w, http.StatusOK, response)
	}
}

func (h *handler) UpdateTaskGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	goalId, err1 := strconv.Atoi(vars["goal_id"])
	taskId, err2 := strconv.Atoi(vars["task_id"])

	if err1 != nil || err2 != nil {
		api.WriteResponse(w, http.StatusBadRequest, apperrors.NewBadRequest("Incorrect request"))
	} else {
		response, appErr := h.service.UpdateTaskGoal(r.Context(), taskId, goalId)
		if appErr != nil {
			api.WriteResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			api.WriteResponse(w, http.StatusOK, response)
		}
	}
}

func (h *handler) DeleteTaskGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["task_id"])

	if err != nil {
		api.WriteResponse(w, http.StatusBadRequest, apperrors.NewBadRequest("Incorrect request"))
	} else {
		response, appErr := h.service.DeleteTaskGoal(r.Context(), taskId)
		if appErr != nil {
			api.WriteResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			api.WriteResponse(w, http.StatusOK, response)
		}
	}
}
