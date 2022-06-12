package task

import (
	"database/sql"
	"doneclub-api/pkg/logging"
)

type Task struct {
	ID         int            `db:"id" json:"id,omitempty,omitempty"`
	UserID     int            `db:"user_id" json:"user_id,omitempty"`
	GoalID     sql.NullInt64  `db:"goal_id" json:"goal_id,omitempty"`
	Status     int            `db:"status" json:"status,omitempty"`
	Title      string         `db:"title" json:"title,omitempty,omitempty"`
	Deadline   sql.NullString `db:"deadline" json:"deadline,omitempty"`
	FinishedAt sql.NullString `db:"finished_at" json:"finished_at,omitempty"`
	CreatedAt  string         `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt  string         `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt  sql.NullString `db:"deleted_at" json:"deleted_at,omitempty"`
}

const (
	banned   = 0
	active   = 1
	inactive = 2
)

func (t *Task) getStatusAsString() string {

	statusDictionary := map[int]string{
		active:   "active",
		inactive: "inactive",
	}
	status, ok := statusDictionary[t.Status]

	if !ok {
		return "undefined"
	}

	return status
}

func GetStatusAsInt(status string) int {
	switch status {
	case "active":
		return active
	case "inactive":
		return inactive
	}
	logging.GetLogger().Error("unknown goal status: " + status)
	return -1
}

func GetAllStatuses() map[string]int {
	statuses := make(map[string]int)

	statuses["banned"] = banned
	statuses["active"] = active
	statuses["inactive"] = inactive

	return statuses
}
