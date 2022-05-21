package user

type CreateUserRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
