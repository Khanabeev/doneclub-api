package user

import (
	"context"
	"doneclub-api/pkg/apperrors"
	"doneclub-api/pkg/logging"
)

type Service interface {
	FindUserById(ctx context.Context, dto *GetUserDTO) (*User, *apperrors.AppError)
	FindUserByEmail(ctx context.Context, userEmail string) (*User, *apperrors.AppError)
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (s *service) FindUserById(ctx context.Context, dto *GetUserDTO) (*User, *apperrors.AppError) {
	return nil, nil
}

func (s *service) FindUserByEmail(ctx context.Context, userEmail string) (*User, *apperrors.AppError) {
	logger := logging.GetLogger()
	user, err := s.storage.GetUserByEmail(userEmail)

	if err != nil {
		logger.Error(err)
		return nil, apperrors.NewUnexpectedError("Unexpected error while validating user email")
	}

	return user, nil
}
