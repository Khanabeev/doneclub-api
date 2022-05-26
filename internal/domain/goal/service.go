package goal

import (
	"context"
	"database/sql"
	"doneclub-api/internal/adapters/api/auth"
	"doneclub-api/pkg/apperrors"
	"doneclub-api/pkg/logging"
)

type Service interface {
	CreateNewGoal(ctx context.Context, dto *RequestCreateGoalDTO) (*ResponseGoalDTO, *apperrors.AppError)
}

type goalService struct {
	storage Storage
}

func NewService(storage Storage) *goalService {
	return &goalService{
		storage: storage,
	}
}

func (g goalService) CreateNewGoal(ctx context.Context, dto *RequestCreateGoalDTO) (*ResponseGoalDTO, *apperrors.AppError) {
	// Validation
	validation := dto.Validate()
	if validation != nil {
		return nil, apperrors.NewValidationError("Validation error", validation)
	}
	logger := logging.GetLogger()
	userClaims, ok := ctx.Value(auth.ContextUserKey).(*auth.UserClaims)
	if !ok {
		logger.Error("Unexpected error while getting user claims in FindUserById method")
	}
	goal := &Goal{
		UserID:      userClaims.ID,
		Status:      1,
		ParentID:    sql.NullString{String: dto.ParentId, Valid: false},
		Title:       dto.Title,
		Description: sql.NullString{String: dto.Description, Valid: false},
		StartDate:   sql.NullString{String: dto.StartDate, Valid: false},
		EndDate:     sql.NullString{String: dto.EndDate, Valid: false},
	}
	newGoal, err := g.storage.CreateGoal(goal)
	if err != nil {
		logger.Error("Can't create a new goal: " + err.Error())
		return nil, apperrors.NewUnexpectedError("Unexpected database error")
	}

	resource := newGoal.GetGoalProfileResource()
	return resource, nil
}
