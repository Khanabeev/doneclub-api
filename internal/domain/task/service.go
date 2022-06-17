package task

import (
	"context"
	"database/sql"
	"doneclub-api/internal/adapters/api/auth"
	"doneclub-api/pkg/apperrors"
	"doneclub-api/pkg/logging"
	"time"
)

type service struct {
	storage Storage
}

func NewService(storage Storage) *service {
	return &service{storage: storage}
}

func (s *service) CreateNewTask(ctx context.Context, dto *RequestCreateTaskDTO) (*ResponseTaskDTO, *apperrors.AppError) {
	// Validation
	validation := dto.Validate()
	if validation != nil {
		return nil, apperrors.NewValidationError("Validation error", validation)
	}
	logger := logging.GetLogger()
	userClaims, ok := ctx.Value(auth.ContextUserKey).(*auth.UserClaims)
	if !ok {
		logger.Error("Unexpected error while getting user claims in CreateNewTask method")
	}
	task := &Task{
		UserID: userClaims.ID,
		Status: GetStatusAsInt(dto.Status),
		Title:  dto.Title,
		Deadline: sql.NullString{
			String: dto.Deadline,
			Valid:  dto.Deadline != "",
		},
	}

	newTask, err := s.storage.CreateTask(task)
	if err != nil {
		logger.Error("Can't create a new task: " + err.Error())
		return nil, apperrors.NewUnexpectedError("Unexpected database error")
	}

	resource := newTask.GetTaskProfileResource()
	return resource, nil
}

func (s *service) GetAllTasks(ctx context.Context, status string) (*ResponseAllTasksDTO, *apperrors.AppError) {
	logger := logging.GetLogger()
	userClaims, ok := ctx.Value(auth.ContextUserKey).(*auth.UserClaims)
	if !ok {
		logger.Error("Unexpected error while getting user claims in FindUserById method")
	}

	var err error
	var tasks []*Task

	allStatuses := GetAllStatuses()
	statusInt, ok := allStatuses[status]

	if !ok {
		tasks, err = s.storage.GetAllTasksByUserId(userClaims.ID)
	} else {
		tasks, err = s.storage.GetAllTasksByUserIdAndStatus(userClaims.ID, statusInt)
	}

	if err == sql.ErrNoRows {
		return nil, apperrors.NewNotFoundError("no tasks found")
	}
	if err != nil {
		return nil, apperrors.NewUnexpectedError(err.Error())
	}

	return GetAllTasksProfileResource(tasks), nil
}

func (s *service) GetAllTasksByGoal(ctx context.Context, status string, goalId int) (*ResponseTasksByGoalDTO, *apperrors.AppError) {
	logger := logging.GetLogger()
	userClaims, ok := ctx.Value(auth.ContextUserKey).(*auth.UserClaims)
	if !ok {
		logger.Error("Unexpected error while getting user claims in FindUserById method")
	}

	var err error
	var tasks []*Task

	allStatuses := GetAllStatuses()
	statusInt, ok := allStatuses[status]

	if !ok {
		tasks, err = s.storage.GetAllTasksByUserIdAndGoalId(userClaims.ID, goalId)
	} else {
		tasks, err = s.storage.GetAllTasksByUserIdAndGoalIdAndStatus(userClaims.ID, goalId, statusInt)
	}

	if err == sql.ErrNoRows {
		return nil, apperrors.NewNotFoundError("no tasks found")
	}
	if err != nil {
		return nil, apperrors.NewUnexpectedError(err.Error())
	}

	return GetAllTasksProfileByGoalResource(tasks, goalId), nil
}

func (s *service) GetTaskById(ctx context.Context, taskId int) (*ResponseTaskDTO, *apperrors.AppError) {
	logger := logging.GetLogger()
	userClaims, ok := ctx.Value(auth.ContextUserKey).(*auth.UserClaims)
	if !ok {
		logger.Error("Unexpected error while getting user claims in GetTaskById method")
	}

	task, err := s.storage.GetTaskById(userClaims.ID, taskId)

	if err == sql.ErrNoRows {
		return nil, apperrors.NewNotFoundError("no tasks found")
	}

	if err != nil {
		return nil, apperrors.NewUnexpectedError(err.Error())
	}

	return task.GetTaskProfileResource(), nil
}

func (s *service) UpdateTask(ctx context.Context, dto *RequestUpdateTaskDTO, taskId int) (*ResponseTaskDTO, *apperrors.AppError) {
	// Validation
	validation := dto.Validate()
	if validation != nil {
		return nil, apperrors.NewValidationError("Validation error", validation)
	}
	logger := logging.GetLogger()
	userClaims, ok := ctx.Value(auth.ContextUserKey).(*auth.UserClaims)
	if !ok {
		logger.Error("Unexpected error while getting user claims in UpdateTask method")
	}

	task := &Task{
		Status: GetStatusAsInt(dto.Status),
		Title:  dto.Title,
		Deadline: sql.NullString{
			String: dto.Deadline,
			Valid:  dto.Deadline != "",
		},
		FinishedAt: sql.NullString{
			String: dto.FinishedAt,
			Valid:  dto.FinishedAt != "",
		},
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	updatedTask, err := s.storage.UpdateTask(userClaims.ID, taskId, task)
	if err != nil {
		_, ok := err.(*apperrors.AppError)
		if ok {
			return nil, apperrors.NewBadRequest(err.Error())
		}

		logger.Error(err.Error())
		return nil, apperrors.NewUnexpectedError("Unexpected database error")
	}

	updatedTask.ID = taskId
	updatedTask.UserID = userClaims.ID
	resource := updatedTask.GetTaskProfileResource()
	return resource, nil
}

func (s *service) DeleteTask(ctx context.Context, taskId int) (*ResponseTaskDTO, *apperrors.AppError) {
	logger := logging.GetLogger()
	userClaims, ok := ctx.Value(auth.ContextUserKey).(*auth.UserClaims)
	if !ok {
		logger.Error("Unexpected error while getting user claims in GetTaskById method")
	}

	task, err := s.storage.DeleteTask(userClaims.ID, taskId)

	if err == sql.ErrNoRows {
		return nil, apperrors.NewNotFoundError("no tasks found")
	}

	if err != nil {
		return nil, apperrors.NewUnexpectedError(err.Error())
	}

	return task.GetTaskProfileResource(), nil
}

func (s *service) UpdateTaskGoal(ctx context.Context, taskId, goalId int) (*ResponseTaskDTO, *apperrors.AppError) {
	logger := logging.GetLogger()
	userClaims, ok := ctx.Value(auth.ContextUserKey).(*auth.UserClaims)
	if !ok {
		logger.Error("Unexpected error while getting user claims in GetTaskById method")
	}

	task, err := s.storage.UpdateTaskGoal(userClaims.ID, taskId, goalId)

	if err == sql.ErrNoRows {
		return nil, apperrors.NewNotFoundError("no tasks found")
	}

	if err != nil {
		logger.Error(err.Error())
		return nil, apperrors.NewBadRequest("bad request")
	}

	return task.GetTaskProfileResource(), nil
}

func (s *service) DeleteTaskGoal(ctx context.Context, taskId int) (*ResponseTaskDTO, *apperrors.AppError) {
	logger := logging.GetLogger()
	userClaims, ok := ctx.Value(auth.ContextUserKey).(*auth.UserClaims)
	if !ok {
		logger.Error("Unexpected error while getting user claims in GetTaskById method")
	}

	task, err := s.storage.DeleteTaskGoal(userClaims.ID, taskId)

	if err == sql.ErrNoRows {
		return nil, apperrors.NewNotFoundError("no tasks found")
	}

	if err != nil {
		logger.Error(err.Error())
		return nil, apperrors.NewBadRequest("incorrect parameters")
	}

	return task.GetTaskProfileResource(), nil
}
