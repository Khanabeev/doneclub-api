package user

import (
	"context"
	"doneclub-api/internal/apperrors"
	"doneclub-api/pkg/logging"
	"doneclub-api/pkg/password_hash"
	"time"
)

type Service interface {
	CreateUser(ctx context.Context, dto *CreateUserRequestDTO) (*CreateUserResponseDTO, *apperrors.AppError)
	GetUserById(ctx context.Context, dto *GetUserDTO) (*User, *apperrors.AppError)
	GetUserByEmail(ctx context.Context, userEmail string) (*User, *apperrors.AppError)
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (s *service) CreateUser(ctx context.Context, dto *CreateUserRequestDTO) (*CreateUserResponseDTO, *apperrors.AppError) {
	logger := logging.GetLogger()

	logger.Info("validate create user request")
	appError := dto.Validate()
	if appError != nil {
		return nil, appError
	}

	logger.Info("Check if email already exist in database")
	user, _ := s.GetUserByEmail(ctx, dto.Email)

	if user != nil {
		return nil, apperrors.NewValidationError("User with this email already exists")
	}

	pass, err := password_hash.HashPassword(dto.Password)
	if err != nil {
		logger.Error(err)
		return nil, apperrors.NewUnexpectedError("Unexpected error while saving the user")
	}

	u := User{
		ID:        0,
		Email:     dto.Email,
		Password:  pass,
		Status:    1,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	logger.Info("create new user")
	user, err = s.storage.CreateUser(&u)
	if err != nil {
		logger.Error(err)
		return nil, apperrors.NewUnexpectedError("can't create user")
	}

	response := &CreateUserResponseDTO{
		Email:  user.Email,
		Status: user.Status,
	}

	return response, nil
}

func (s *service) GetUserById(ctx context.Context, dto *GetUserDTO) (*User, *apperrors.AppError) {
	return nil, nil
}

func (s *service) GetUserByEmail(ctx context.Context, userEmail string) (*User, *apperrors.AppError) {
	logger := logging.GetLogger()
	user, err := s.storage.GetUserByEmail(userEmail)

	if err != nil {
		logger.Error(err)
		return nil, apperrors.NewUnexpectedError("Unexpected error while validating user email")
	}

	return user, nil
}
