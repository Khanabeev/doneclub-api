package category

import "database/sql"

type Category struct {
	ID          int            `db:"id" json:"id,omitempty"`
	UserID      int            `db:"user_id" json:"user_id,omitempty"`
	Title       string         `db:"title" json:"title,omitempty"`
	Description sql.NullString `db:"description" json:"description"`
	CreatedAt   string         `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt   string         `db:"updated_at" json:"updated_at,omitempty"`
}
