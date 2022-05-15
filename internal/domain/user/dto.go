package user

import "doneclub-api/internal/apperrors"

type CreateUserRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c CreateUserRequestDTO) Validate() *apperrors.AppError {
	if len(c.Email) > 20 {
		return apperrors.NewValidationError("The length of email is more than 20 symbols")
	}

	return nil
}

type CreateUserResponseDTO struct {
	Email  string `json:"email"`
	Status int    `json:"status"`
}

type GetUserDTO struct {
	ID     int    `json:"id"`
	Email  string `json:"email"`
	Status int    `json:"status"`
}
