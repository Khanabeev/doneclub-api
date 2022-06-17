package goal

import (
	"context"
	"doneclub-api/internal/domain/goal"
	"doneclub-api/pkg/apperrors"
)

type Service interface {
	CreateNewGoal(ctx context.Context, dto *goal.RequestCreateGoalDTO) (*goal.ResponseGoalDTO, *apperrors.AppError)
	UpdateGoal(ctx context.Context, dto *goal.RequestUpdateGoalDTO, goalId int) (*goal.ResponseGoalDTO, *apperrors.AppError)
	GetGoal(ctx context.Context, goalId int) (*goal.ResponseGoalDTO, *apperrors.AppError)
	GetAllGoals(ctx context.Context, status string) (*goal.ResponseAllGoalsDTO, *apperrors.AppError)
	DeleteGoal(ctx context.Context, goalId int) (*goal.ProfileGoalDeleted, *apperrors.AppError)
	UpdateGoalParentId(ctx context.Context, goalId, parentId int) (*goal.ResponseGoalDTO, *apperrors.AppError)
	DeleteGoalParentId(ctx context.Context, goalId int) (*goal.ResponseGoalDTO, *apperrors.AppError)
}
