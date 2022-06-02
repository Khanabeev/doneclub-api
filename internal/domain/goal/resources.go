package goal

import (
	"doneclub-api/internal/domain/user"
)

func (g *Goal) GetGoalProfileResource() *ResponseGoalDTO {
	return &ResponseGoalDTO{
		Goal: &ProfileGoal{
			ID: g.ID,
			User: user.ProfileUser{
				ID: g.UserID,
			},
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
