package user

import (
	"context"
	"doneclub-api/internal/adapters/api/auth"
	"doneclub-api/pkg/apperrors"
	"doneclub-api/pkg/logging"
)

type Service interface {
	FindUserById(ctx context.Context) (*ResponseUserDTO, *apperrors.AppError)
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (s *service) FindUserById(ctx context.Context) (*ResponseUserDTO, *apperrors.AppError) {
	logger := logging.GetLogger()
	userClaims, ok := ctx.Value(auth.ContextUserKey).(*auth.UserClaims)
	if !ok {
		logger.Error("Unexpected error while getting user claims in FindUserById method")
	}

	u, err := s.storage.GetUserById(userClaims.ID)
	if err != nil {
		logger.Error(err)
		return nil, apperrors.NewUnexpectedError("Unexpected error while validating user email")
	}
	response := u.GetUserProfileResource()

	return response, nil
}
