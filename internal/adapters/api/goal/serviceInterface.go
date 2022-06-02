package goal

import (
	"context"
	"doneclub-api/internal/domain/goal"
	"doneclub-api/pkg/apperrors"
)

type Service interface {
	CreateNewGoal(ctx context.Context, dto *goal.RequestCreateGoalDTO) (*goal.ResponseGoalDTO, *apperrors.AppError)
	UpdateGoal(ctx context.Context, dto *goal.RequestUpdateGoalDTO, goalId string) (*goal.ResponseGoalDTO, *apperrors.AppError)
}
