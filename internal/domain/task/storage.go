package task

type Storage interface {
	CreateTask(task *Task) (*Task, error)
	GetAllTasksByUserId(userId int) ([]*Task, error)
	GetAllTasksByUserIdAndStatus(userId, status int) ([]*Task, error)
	GetAllTasksByUserIdAndGoalId(userId, goalId int) ([]*Task, error)
	GetAllTasksByUserIdAndGoalIdAndStatus(userId, goalId, status int) ([]*Task, error)
	GetTaskById(userId, taskId int) (*Task, error)
	UpdateTask(userId, taskId int, task *Task) (*Task, error)
	DeleteTask(userId, taskId int) (*Task, error)
	UpdateTaskGoal(userId, taskId, goalId int) (*Task, error)
	DeleteTaskGoal(userId, taskId int) (*Task, error)
}
