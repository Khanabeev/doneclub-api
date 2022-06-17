package task

func (t *Task) GetTaskProfileResource() *ResponseTaskDTO {
	return &ResponseTaskDTO{
		Task: &ProfileTask{
			ID:         t.ID,
			UserID:     t.UserID,
			GoalID:     t.GoalID.Int64,
			Title:      t.Title,
			Deadline:   t.Deadline.String,
			FinishedAt: t.FinishedAt.String,
			CreatedAt:  t.CreatedAt,
			Status:     t.getStatusAsString(),
		},
	}
}

func GetAllTasksProfileResource(tasks []*Task) *ResponseAllTasksDTO {
	var allTasks []*ProfileTask
	for _, task := range tasks {
		allTasks = append(allTasks, task.GetTaskProfileResource().Task)
	}
	return &ResponseAllTasksDTO{
		Tasks: allTasks,
	}
}

func GetAllTasksProfileByGoalResource(tasks []*Task, goalId int) *ResponseTasksByGoalDTO {
	var allTasks []*ProfileTask
	for _, task := range tasks {
		allTasks = append(allTasks, task.GetTaskProfileResource().Task)
	}
	return &ResponseTasksByGoalDTO{
		GoalId: goalId,
		Tasks:  allTasks,
	}
}
