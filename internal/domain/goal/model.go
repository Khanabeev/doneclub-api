package goal

import (
	"database/sql"
)

type Goal struct {
	ID          int            `db:"id" json:"id,omitempty,omitempty"`
	UserID      int            `db:"user_id" json:"user_id,omitempty"`
	Status      int            `db:"status" json:"status,omitempty"`
	ParentID    sql.NullInt64  `db:"parent_id" json:"parent_id,omitempty"`
	Title       string         `db:"title" json:"title,omitempty,omitempty"`
	Description sql.NullString `db:"description" json:"description,omitempty"`
	StartDate   sql.NullString `db:"start_date" json:"start_date,omitempty"`
	EndDate     sql.NullString `db:"end_date" json:"end_date,omitempty"`
	CreatedAt   string         `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt   string         `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt   sql.NullString `db:"deleted_at" json:"deleted_at,omitempty"`
}

const (
	active   = 1
	inactive = 2
)

func (u *Goal) getStatusAsString() string {

	statusDictionary := map[int]string{
		active:   "active",
		inactive: "inactive",
	}
	status, ok := statusDictionary[u.Status]

	if !ok {
		return "undefined"
	}

	return status
}
