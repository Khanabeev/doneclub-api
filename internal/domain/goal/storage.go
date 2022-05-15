package goal

import "doneclub-api/internal/domain/user"

type Storage interface {
	GetAllGoalsByUser(user *user.User, limit int, offset int) []*Goal
	GetGoalById(goalId int) []*Goal
}
