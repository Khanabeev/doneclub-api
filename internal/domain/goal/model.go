package goal

import (
	"database/sql"
	"doneclub-api/internal/domain/user"
)

type Goal struct {
	ID          int            `json:"id,omitempty,omitempty"`
	User        user.User      `json:"user,omitempty"`
	Status      int            `json:"status,omitempty,omitempty"`
	ParentID    sql.NullInt32  `json:"parent_id,omitempty"`
	Title       string         `json:"title,omitempty,omitempty"`
	Description sql.NullString `json:"description,omitempty"`
	StartDate   sql.NullString `json:"start_date,omitempty"`
	EndDate     sql.NullString `json:"end_date,omitempty"`
	CreatedAt   string         `json:"created_at,omitempty"`
	UpdatedAt   string         `json:"updated_at,omitempty"`
	DeletedAt   string         `json:"deleted_at,omitempty"`
}
