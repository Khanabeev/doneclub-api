package task

import (
	"doneclub-api/internal/domain/task"
	"doneclub-api/pkg/logging"
	"github.com/jmoiron/sqlx"
)

type storage struct {
	client *sqlx.DB
}

func NewStorage(client *sqlx.DB) task.Storage {
	return &storage{
		client: client,
	}
}

func (s *storage) CreateTask(t *task.Task) (*task.Task, error) {
	logger := logging.GetLogger()
	query := "INSERT INTO tasks (user_id, status, title, deadline) VALUES(?, ?, ?, ?)"
	result, err := s.client.Exec(query, t.UserID, t.Status, t.Title, t.Deadline)
	if err != nil {
		logger.Errorf("unexpected error during insert new task: %s", err.Error())
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Errorf("unexpected error while getting last inserted task id: %s", err.Error())
		return nil, err
	}

	t.ID = int(id)
	return t, nil
}

func (s *storage) GetAllTasksByUserId(userId int) ([]*task.Task, error) {
	query := `SELECT * 
				FROM tasks
				WHERE user_id = ?
				  AND deleted_at IS NULL`
	var tasks []*task.Task
	err := s.client.Select(&tasks, query, userId)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *storage) GetAllTasksByUserIdAndStatus(userId, status int) ([]*task.Task, error) {
	query := `SELECT * 
				FROM tasks
				WHERE user_id = ?
				  AND status = ?
				  AND deleted_at IS NULL`
	var tasks []*task.Task
	err := s.client.Select(&tasks, query, userId, status)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *storage) GetAllTasksByUserIdAndGoalId(userId, goalId int) ([]*task.Task, error) {
	query := `SELECT * 
				FROM tasks
				WHERE user_id = ?
				  AND goal_id = ?
				  AND deleted_at IS NULL`
	var tasks []*task.Task
	err := s.client.Select(&tasks, query, userId, goalId)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
func (s *storage) GetAllTasksByUserIdAndGoalIdAndStatus(userId, goalId, status int) ([]*task.Task, error) {
	query := `SELECT * 
				FROM tasks
				WHERE user_id = ?
				  AND goal_id = ?
				  AND status = ?
				  AND deleted_at IS NULL`
	var tasks []*task.Task
	err := s.client.Select(&tasks, query, userId, goalId, status)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
