package goal

import (
	"database/sql"
	"doneclub-api/internal/domain/user"
	"doneclub-api/pkg/input_validator"
)

type RequestCreateGoalDTO struct {
	ParentId    sql.NullInt64 `json:"parent_id,omitempty" validate:"omitempty,numeric"`
	Status      string        `json:"status" validate:"required,oneof=active inactive"`
	Title       string        `json:"title,omitempty" validate:"required,lte=255,gte=2"`
	Description string        `json:"description,omitempty" validate:"lte=1000,gte=2"`
	StartDate   string        `json:"start_date,omitempty" validate:"omitempty,datetime=2006-01-02 15:04:05"`
	EndDate     string        `json:"end_date,omitempty" validate:"omitempty,datetime=2006-01-02 15:04:05"`
}

func (r *RequestCreateGoalDTO) Validate() []string {
	return input_validator.NewInputValidator().Validate(r)
}

type RequestUpdateGoalDTO struct {
	ParentId    sql.NullInt64 `json:"parent_id,omitempty" validate:"omitempty,numeric"`
	Status      string        `json:"status" validate:"required"`
	Title       string        `json:"title,omitempty" validate:"required,lte=255,gte=2"`
	Description string        `json:"description,omitempty" validate:"omitempty,lte=1000,gte=2"`
	StartDate   string        `json:"start_date,omitempty" validate:"omitempty,datetime=2006-01-02 15:04:05"`
	EndDate     string        `json:"end_date,omitempty" validate:"omitempty,datetime=2006-01-02 15:04:05"`
}

func (r *RequestUpdateGoalDTO) Validate() []string {
	return input_validator.NewInputValidator().Validate(r)
}

type ResponseGoalDTO struct {
	Goal *ProfileGoal `json:"goal"`
}

type ResponseAllGoalsDTO struct {
	Goals []*ProfileGoal `json:"goals"`
}

type ProfileGoal struct {
	ID          int              `json:"id,omitempty"`
	User        user.ProfileUser `json:"user,omitempty"`
	ParentId    int64            `json:"parent_id,omitempty"`
	Title       string           `json:"title,omitempty"`
	Description string           `json:"description,omitempty"`
	StartDate   string           `json:"start_date,omitempty"`
	EndDate     string           `json:"end_date,omitempty"`
	CreatedAt   string           `json:"created_at,omitempty"`
	UpdatedAt   string           `json:"updated_at,omitempty"`
	Status      string           `json:"status,omitempty"`
}

type ProfileGoalDeleted struct {
	ID     int `json:"id,omitempty"`
	UserId int `json:"user_id,omitempty"`
}
