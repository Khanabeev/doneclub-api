package user

func (u *User) GetUserProfileResource() *ResponseUserDTO {
	return &ResponseUserDTO{
		User: &ProfileUser{
			ID:        u.ID,
			Email:     u.Email,
			Status:    u.getStatusAsString(),
			CreatedAt: u.CreatedAt,
		},
	}
}
