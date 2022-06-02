package goal

import (
	"database/sql"
)

type Goal struct {
	ID          int            `json:"id,omitempty,omitempty"`
	UserID      int            `json:"user,omitempty"`
	Status      int            `json:"status,omitempty,omitempty"`
	ParentID    sql.NullInt64  `json:"parent_id,omitempty"`
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
)

func (u *Goal) getStatusAsString() string {
	var status string
	switch u.Status {
	case active:
		status = "active"
	case inactive:
		status = "inactive"
	}
	return status
}

func GetStatusAsInt(status string) int {
	switch status {
	case "active":
		return 1
	case "inactive":
		return 2
	}
	return 0
}
