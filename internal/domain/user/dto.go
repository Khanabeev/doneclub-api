package user

type ResponseUserDTO struct {
	User *ProfileUser `json:"user"`
}

type ProfileUser struct {
	ID        int    `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Status    string `json:"status,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
