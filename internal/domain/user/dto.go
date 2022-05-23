package user

type ResponseGetUserProfileDTO struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"`
}
