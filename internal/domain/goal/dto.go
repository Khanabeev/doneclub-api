package goal

import (
	"doneclub-api/internal/domain/user"
	"doneclub-api/pkg/input_validator"
)

type RequestCreateGoalDTO struct {
	ParentId    string `json:"parent_id,omitempty" validate:"omitempty,numeric"`
	Title       string `json:"title,omitempty" validate:"required,lte=255,gte=2"`
	Description string `json:"description,omitempty" validate:"lte=1000,gte=2"`
	StartDate   string `json:"start_date,omitempty" validate:"omitempty,datetime=2006-01-02 01:01:01"`
	EndDate     string `json:"end_date,omitempty" validate:"omitempty,datetime=2006-01-02 01:01:01"`
}

func (r *RequestCreateGoalDTO) Validate() []string {
	return input_validator.NewInputValidator().Validate(r)
}

type ResponseGoalDTO struct {
	Goal *ProfileGoal `json:"goal"`
}
type ProfileGoal struct {
	ID          string           `json:"id,omitempty"`
	User        user.ProfileUser `json:"user"`
	ParentId    string           `json:"parent_id,omitempty"`
	Title       string           `json:"title,omitempty"`
	Description string           `json:"description,omitempty"`
	StartDate   string           `json:"start_date,omitempty"`
	EndDate     string           `json:"end_date,omitempty"`
	CreatedAt   string           `json:"created_at,omitempty"`
	UpdatedAt   string           `json:"updated_at,omitempty"`
	Status      string           `json:"status,omitempty"`
}

type ResponseGetOneGoalDTO struct {
	ParentId    string `json:"parent_id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	StartDate   string `json:"start_date,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	Status      string `json:"status,omitempty"`
}
