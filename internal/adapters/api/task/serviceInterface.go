package task

import (
	"context"
	"doneclub-api/internal/domain/task"
	"doneclub-api/pkg/apperrors"
)

type Service interface {
	CreateNewTask(ctx context.Context, dto *task.RequestCreateTaskDTO) (*task.ResponseTaskDTO, *apperrors.AppError)
	GetAllTasks(ctx context.Context, status string) (*task.ResponseAllTasksDTO, *apperrors.AppError)
	GetAllTasksByGoal(ctx context.Context, status string, goalId int) (*task.ResponseTasksByGoalDTO, *apperrors.AppError)
	GetTaskById(ctx context.Context, taskId int) (*task.ResponseTaskDTO, *apperrors.AppError)
	UpdateTask(ctx context.Context, dto *task.RequestUpdateTaskDTO, taskId int) (*task.ResponseTaskDTO, *apperrors.AppError)
	DeleteTask(ctx context.Context, taskId int) (*task.ResponseTaskDTO, *apperrors.AppError)
	UpdateTaskGoal(ctx context.Context, taskId, goalId int) (*task.ResponseTaskDTO, *apperrors.AppError)
	DeleteTaskGoal(ctx context.Context, taskId int) (*task.ResponseTaskDTO, *apperrors.AppError)
}
