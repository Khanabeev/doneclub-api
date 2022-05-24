package user

type ResponseGetUserProfileDTO struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}
