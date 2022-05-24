package user

type User struct {
	ID        int    `db:"id" json:"id,omitempty"`
	Email     string `db:"email" json:"email,omitempty"`
	Password  string `db:"password" json:"-"`
	Status    int    `db:"status" json:"status,omitempty"`
	Role      string `db:"role" json:"role,omitempty"`
	CreatedAt string `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt string `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt string `db:"deleted_at" json:"deleted_at,omitempty"`
}

const (
	active   = 1
	inactive = 2
	banned   = 3
)

func (u *User) ToDtoUserProfile() *ResponseGetUserProfileDTO {
	return &ResponseGetUserProfileDTO{
		ID:        u.ID,
		Email:     u.Email,
		Status:    u.getStatusAsString(),
		CreatedAt: u.CreatedAt,
	}
}

func (u *User) getStatusAsString() string {
	var status string
	switch u.Status {
	case active:
		status = "active"
	case inactive:
		status = "inactive"
	case banned:
		status = "banned"
	}
	return status
}
