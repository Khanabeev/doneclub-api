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

func (s *service) CreateNewGoal(ctx context.Context, dto *RequestCreateGoalDTO) (*ResponseGoalDTO, *apperrors.AppError) {
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
		ParentID:    sql.NullInt64{Int64: dto.ParentId, Valid: dto.ParentId != 0},
		Title:       dto.Title,
		Description: sql.NullString{String: dto.Description, Valid: dto.Description != ""},
		StartDate:   sql.NullString{String: dto.StartDate, Valid: dto.StartDate != ""},
		EndDate:     sql.NullString{String: dto.EndDate, Valid: dto.EndDate != ""},
	}
	newGoal, err := s.storage.CreateGoal(goal)
	if err != nil {
		logger.Error("Can't create a new goal: " + err.Error())
		return nil, apperrors.NewUnexpectedError("Unexpected database error")
	}

	resource := newGoal.GetGoalProfileResource()
	return resource, nil
}

func (s *service) UpdateGoal(ctx context.Context, dto *RequestUpdateGoalDTO, goalId int) (*ResponseGoalDTO, *apperrors.AppError) {
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
		ParentID:    sql.NullInt64{Int64: dto.ParentId, Valid: dto.ParentId != 0},
		Title:       dto.Title,
		Description: sql.NullString{String: dto.Description, Valid: dto.Description != ""},
		StartDate:   sql.NullString{String: dto.StartDate, Valid: dto.StartDate != ""},
		EndDate:     sql.NullString{String: dto.EndDate, Valid: dto.EndDate != ""},
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}
	updatedGoal, err := s.storage.UpdateGoal(goal, goalId)
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

func (s *service) GetGoal(ctx context.Context, goalId int) (*ResponseGoalDTO, *apperrors.AppError) {
	logger := logging.GetLogger()
	userClaims, ok := ctx.Value(auth.ContextUserKey).(*auth.UserClaims)
	if !ok {
		logger.Error("Unexpected error while getting user claims in FindUserById method")
		return nil, apperrors.NewUnauthorizedError("unexpected error")
	}

	goal, err := s.storage.GetGoalById(userClaims.ID, goalId)
	if err == sql.ErrNoRows {
		return nil, apperrors.NewNotFoundError("no goals found")
	}
	if err != nil {
		return nil, apperrors.NewUnexpectedError(err.Error())
	}

	resource := goal.GetGoalProfileResource()
	return resource, nil
}

func (s *service) GetAllGoals(ctx context.Context, status string) (*ResponseAllGoalsDTO, *apperrors.AppError) {
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
		goals, err = s.storage.GetAllGoalsByUserId(userClaims.ID)
	} else {
		goals, err = s.storage.GetAllGoalsByUserIdAndStatus(userClaims.ID, statusInt)
	}

	if err == sql.ErrNoRows {
		return nil, apperrors.NewNotFoundError("no goals found")
	}
	if err != nil {
		return nil, apperrors.NewUnexpectedError(err.Error())
	}

	return GetAllGoalsProfileResource(goals), nil
}

func (s *service) DeleteGoal(ctx context.Context, goalId int) (*ProfileGoalDeleted, *apperrors.AppError) {
	logger := logging.GetLogger()
	userClaims, ok := ctx.Value(auth.ContextUserKey).(*auth.UserClaims)
	if !ok {
		logger.Error("Unexpected error while getting user claims in DeleteGoal method")
		return nil, apperrors.NewUnauthorizedError("unexpected error")
	}

	err := s.storage.DeleteGoalById(userClaims.ID, goalId)

	if err != nil {
		return nil, apperrors.NewNotFoundError("no goals deleted")
	}

	return DeletedGoalResource(goalId, userClaims.ID), nil
}

func (s *service) UpdateGoalParentId(ctx context.Context, goalId, parentId int) (*ResponseGoalDTO, *apperrors.AppError) {
	logger := logging.GetLogger()
	userClaims, ok := ctx.Value(auth.ContextUserKey).(*auth.UserClaims)
	if !ok {
		logger.Error("Unexpected error while getting user claims in FindUserById method")
		return nil, apperrors.NewUnauthorizedError("unexpected error")
	}

	goal, err := s.storage.UpdateGoalParentId(userClaims.ID, goalId, parentId)
	if err == sql.ErrNoRows {
		return nil, apperrors.NewNotFoundError("no goals found")
	}
	if err != nil {
		return nil, apperrors.NewUnexpectedError(err.Error())
	}

	resource := goal.GetGoalProfileResource()
	return resource, nil
}

func (s *service) DeleteGoalParentId(ctx context.Context, goalId int) (*ResponseGoalDTO, *apperrors.AppError) {
	logger := logging.GetLogger()
	userClaims, ok := ctx.Value(auth.ContextUserKey).(*auth.UserClaims)
	if !ok {
		logger.Error("Unexpected error while getting user claims in FindUserById method")
		return nil, apperrors.NewUnauthorizedError("unexpected error")
	}

	goal, err := s.storage.DeleteGoalParentId(userClaims.ID, goalId)
	if err == sql.ErrNoRows {
		return nil, apperrors.NewNotFoundError("no goals found")
	}
	if err != nil {
		return nil, apperrors.NewUnexpectedError(err.Error())
	}

	resource := goal.GetGoalProfileResource()
	return resource, nil
}
