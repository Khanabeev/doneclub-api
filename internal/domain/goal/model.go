package goal

import (
	"database/sql"
)

type Goal struct {
	ID          int            `json:"id,omitempty,omitempty"`
	UserID      int            `json:"user,omitempty"`
	Status      int            `json:"status,omitempty,omitempty"`
	ParentID    sql.NullString `json:"parent_id,omitempty"`
	Title       string         `json:"title,omitempty,omitempty"`
	Description sql.NullString `json:"description,omitempty"`
	StartDate   sql.NullString `json:"start_date,omitempty"`
	EndDate     sql.NullString `json:"end_date,omitempty"`
	CreatedAt   string         `json:"created_at,omitempty"`
	UpdatedAt   string         `json:"updated_at,omitempty"`
	DeletedAt   string         `json:"deleted_at,omitempty"`
}

const (
	active   = 1
	inactive = 2
	banned   = 3
)

func (u *Goal) getStatusAsString() string {
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
