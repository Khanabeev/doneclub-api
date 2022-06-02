package goal

import "doneclub-api/internal/domain/user"

type Storage interface {
	GetAllGoalsByUser(user *user.User, limit int, offset int) ([]*Goal, error)
	GetGoalById(userId, goalId int) (*Goal, error)
	CreateGoal(goal *Goal) (*Goal, error)
	UpdateGoal(goal *Goal, goalId string) (*Goal, error)
}
