package user

import (
	"context"
	"doneclub-api/internal/domain/user"
	"doneclub-api/pkg/apperrors"
)

type Service interface {
	FindUserById(ctx context.Context) (*user.ResponseUserDTO, *apperrors.AppError)
}
