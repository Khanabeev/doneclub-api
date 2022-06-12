package task

import (
	"doneclub-api/pkg/input_validator"
)

type RequestCreateTaskDTO struct {
	Status   string `json:"status" validate:"required,oneof=active inactive"`
	Title    string `json:"title,omitempty" validate:"required,lte=255,gte=2"`
	Deadline string `json:"deadline,omitempty" validate:"omitempty,datetime=2006-01-02 15:04:05"`
}

func (r *RequestCreateTaskDTO) Validate() []string {
	return input_validator.NewInputValidator().Validate(r)
}

type ResponseTaskDTO struct {
	Task *ProfileTask `json:"task"`
}

type ResponseAllTasksDTO struct {
	Tasks []*ProfileTask `json:"tasks,omitempty"`
}

type ResponseTasksByGoalDTO struct {
	GoalId int            `json:"goal_id"`
	Tasks  []*ProfileTask `json:"tasks,omitempty"`
}

type ProfileTask struct {
	ID        int    `json:"id,omitempty"`
	UserID    int    `json:"user_id"`
	GoalID    int64  `json:"goal_id,omitempty"`
	Title     string `json:"title,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Status    string `json:"status,omitempty"`
	Deadline  string `json:"deadline,omitempty"`
}
