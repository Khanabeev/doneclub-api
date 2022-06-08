package goal

type Storage interface {
	GetAllGoalsByUserIdAndStatus(userId int, status int) ([]*Goal, error)
	GetAllGoalsByUserId(userId int) ([]*Goal, error)
	GetGoalById(userId int, goalId int) (*Goal, error)
	DeleteGoalById(userId int, goalId int) error
	CreateGoal(goal *Goal) (*Goal, error)
	UpdateGoal(goal *Goal, goalId int) (*Goal, error)
}
