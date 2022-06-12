package task

type Storage interface {
	CreateTask(task *Task) (*Task, error)
	GetAllTasksByUserId(userId int) ([]*Task, error)
	GetAllTasksByUserIdAndStatus(userId, status int) ([]*Task, error)
	GetAllTasksByUserIdAndGoalId(userId, goalId int) ([]*Task, error)
	GetAllTasksByUserIdAndGoalIdAndStatus(userId, goalId, status int) ([]*Task, error)
}
