package goal

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

func (g service) UpdateGoal(ctx context.Context, dto *RequestUpdateGoalDTO, goalId int) (*ResponseGoalDTO, *apperrors.AppError) {
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

	updatedGoal.ID = goalId
	resource := updatedGoal.GetGoalProfileResource()
	return resource, nil
}

func (g service) GetGoal(ctx context.Context, goalId int) (*ResponseGoalDTO, *apperrors.AppError) {
	logger := logging.GetLogger()
	userClaims, ok := ctx.Value(auth.ContextUserKey).(*auth.UserClaims)
	if !ok {
		logger.Error("Unexpected error while getting user claims in FindUserById method")
		return nil, apperrors.NewUnauthorizedError("unexpected error")
	}

	goal, err := g.storage.GetGoalById(userClaims.ID, goalId)
	if err == sql.ErrNoRows {
		return nil, apperrors.NewNotFoundError("no goals found")
	}
	if err != nil {
		return nil, apperrors.NewUnexpectedError(err.Error())
	}

	resource := goal.GetGoalProfileResource()
	return resource, nil
}

func (g service) GetAllGoals(ctx context.Context, status string) (*ResponseAllGoalsDTO, *apperrors.AppError) {
	logger := logging.GetLogger()
	userClaims, ok := ctx.Value(auth.ContextUserKey).(*auth.UserClaims)
	if !ok {
		logger.Error("Unexpected error while getting user claims in FindUserById method")
	}

	var err error
	var goals []*Goal

	allStatuses := GetAllStatuses()
	statusInt, ok := allStatuses[status]

	if !ok {
		goals, err = g.storage.GetAllGoalsByUserId(userClaims.ID)
	} else {
		goals, err = g.storage.GetAllGoalsByUserIdAndStatus(userClaims.ID, statusInt)
	}

	if err == sql.ErrNoRows {
		return nil, apperrors.NewNotFoundError("no goals found")
	}
	if err != nil {
		return nil, apperrors.NewUnexpectedError(err.Error())
	}

	return GetAllGoalsProfileResource(goals), nil
}

func (g service) DeleteGoal(ctx context.Context, goalId int) (*ProfileGoalDeleted, *apperrors.AppError) {
	logger := logging.GetLogger()
	userClaims, ok := ctx.Value(auth.ContextUserKey).(*auth.UserClaims)
	if !ok {
		logger.Error("Unexpected error while getting user claims in DeleteGoal method")
		return nil, apperrors.NewUnauthorizedError("unexpected error")
	}

	err := g.storage.DeleteGoalById(userClaims.ID, goalId)

	if err != nil {
		return nil, apperrors.NewNotFoundError("no goals deleted")
	}

	return DeletedGoalResource(goalId, userClaims.ID), nil
}

func GetStatusAsInt(status string) int {
	switch status {
	case "active":
		return 1
	case "inactive":
		return 2
	}
	logging.GetLogger().Error("unknown goal status: " + status)
	return 0
}

func GetAllStatuses() map[string]int {
	statuses := make(map[string]int)

	statuses["banned"] = 0
	statuses["active"] = 1
	statuses["inactive"] = 2

	return statuses

}
