package goal

import "doneclub-api/internal/domain/user"

type Goal struct {
	ID          int       `json:"id"`
	User        user.User `json:"user,omitempty"`
	Status      int       `json:"status,omitempty"`
	ParentID    int       `json:"parent_id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	StartDate   string    `json:"start_date,omitempty"`
	EndDate     string    `json:"end_date,omitempty"`
}
