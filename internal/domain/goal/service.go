package goal

import (
	"context"
	"database/sql"
	"doneclub-api/internal/adapters/api/auth"
	"doneclub-api/pkg/apperrors"
	"doneclub-api/pkg/logging"
	"strconv"
	"time"
)

type service struct {
	storage Storage
}

func NewService(storage Storage) *service {
	return &service{
		storage: storage,
	}
}

func (g service) CreateNewGoal(ctx context.Context, dto *RequestCreateGoalDTO) (*ResponseGoalDTO, *apperrors.AppError) {
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
		Status:      GetStatusAsInt(dto.Status),
		ParentID:    dto.ParentId,
		Title:       dto.Title,
		Description: dto.Description,
		StartDate:   dto.StartDate,
		EndDate:     dto.EndDate,
	}
	newGoal, err := g.storage.CreateGoal(goal)
	if err != nil {
		logger.Error("Can't create a new goal: " + err.Error())
		return nil, apperrors.NewUnexpectedError("Unexpected database error")
	}

	resource := newGoal.GetGoalProfileResource()
	return resource, nil
}

func (g service) UpdateGoal(ctx context.Context, dto *RequestUpdateGoalDTO, goalId string) (*ResponseGoalDTO, *apperrors.AppError) {
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
		Status:      GetStatusAsInt(dto.Status),
		ParentID:    dto.ParentId,
		Title:       dto.Title,
		Description: sql.NullString{String: dto.Description, Valid: dto.Description != ""},
		StartDate:   sql.NullString{String: dto.StartDate, Valid: dto.StartDate != ""},
		EndDate:     sql.NullString{String: dto.EndDate, Valid: dto.EndDate != ""},
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}
	updatedGoal, err := g.storage.UpdateGoal(goal, goalId)
	if err != nil {
		_, ok := err.(*apperrors.AppError)
		if ok {
			return nil, apperrors.NewBadRequest(err.Error())
		}

		logger.Error(err.Error())
		return nil, apperrors.NewUnexpectedError("Unexpected database error")
	}

	updatedGoal.ID, _ = strconv.Atoi(goalId)
	resource := updatedGoal.GetGoalProfileResource()
	return resource, nil
}
