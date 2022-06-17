package goal

type Storage interface {
	GetAllGoalsByUserIdAndStatus(userId int, status int) ([]*Goal, error)
	GetAllGoalsByUserId(userId int) ([]*Goal, error)
	GetGoalById(userId, goalId int) (*Goal, error)
	DeleteGoalById(userId, goalId int) error
	CreateGoal(goal *Goal) (*Goal, error)
	UpdateGoal(goal *Goal, goalId int) (*Goal, error)
	UpdateGoalParentId(userId, goalId, parentId int) (*Goal, error)
	DeleteGoalParentId(userId, goalId int) (*Goal, error)
}
