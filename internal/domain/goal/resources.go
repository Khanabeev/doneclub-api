package goal

func (g *Goal) GetGoalProfileResource() *ResponseGoalDTO {
	return &ResponseGoalDTO{
		Goal: &ProfileGoal{
			ID:          g.ID,
			UserId:      g.UserID,
			ParentId:    g.ParentID.Int64,
			Title:       g.Title,
			Description: g.Description.String,
			StartDate:   g.StartDate.String,
			EndDate:     g.EndDate.String,
			CreatedAt:   g.CreatedAt,
			UpdatedAt:   g.UpdatedAt,
			Status:      g.getStatusAsString(),
		},
	}
}

func GetAllGoalsProfileResource(goals []*Goal) *ResponseAllGoalsDTO {
	var allGoals []*ProfileGoal
	for _, goal := range goals {
		allGoals = append(allGoals, goal.GetGoalProfileResource().Goal)
	}
	return &ResponseAllGoalsDTO{
		Goals: allGoals,
	}
}

func DeletedGoalResource(goalId int, userId int) *ProfileGoalDeleted {
	return &ProfileGoalDeleted{
		ID:     goalId,
		UserId: userId,
	}
}
