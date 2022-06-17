package task

import (
	"doneclub-api/internal/domain/task"
	"doneclub-api/pkg/apperrors"
	"doneclub-api/pkg/logging"
	"errors"
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

func (s *storage) GetTaskById(userId, taskId int) (*task.Task, error) {
	query := `SELECT * 
				FROM tasks
				WHERE user_id = ?
				  AND id = ?
				  AND deleted_at IS NULL`
	var selectedTask task.Task
	err := s.client.Get(&selectedTask, query, userId, taskId)
	if err != nil {
		return nil, err
	}

	return &selectedTask, nil
}

func (s *storage) UpdateTask(userId, taskId int, task *task.Task) (*task.Task, error) {
	logger := logging.GetLogger()
	query := `UPDATE tasks 
				SET 
					status = ?,
					title = ?,
					deadline = ?,
					finished_at = ?,
				    updated_at = ?
				WHERE user_id = ? 
				  AND id = ? 
				  AND deleted_at IS NULL`

	result, err := s.client.Exec(
		query,
		task.Status,
		task.Title,
		task.Deadline,
		task.FinishedAt,
		task.UpdatedAt,
		userId,
		taskId,
	)
	if err != nil {
		logger.Errorf("unexpected error during update a task: %s", err.Error())
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		logger.Errorf("unexpected error while getting rows affected: %s", err.Error())
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, apperrors.NewBadRequest("incorrect task parameters")
	}

	task, err = s.GetTaskById(userId, taskId)
	if err != nil {
		logger.Errorf("unexpected error while getting task by id: %s", err.Error())
		return nil, err
	}
	return task, nil

}

func (s *storage) DeleteTask(userId, taskId int) (*task.Task, error) {

	logger := logging.GetLogger()
	query := `UPDATE tasks 
				SET 
					deleted_at = CURRENT_TIMESTAMP
				WHERE user_id = ? 
				  AND id = ? 
				  AND deleted_at IS NULL`

	taskToDelete, err := s.GetTaskById(userId, taskId)
	if err != nil {
		return nil, err
	}

	result, err := s.client.Exec(
		query,
		userId,
		taskId,
	)
	if err != nil {
		logger.Errorf("unexpected error during deleting a task: %s", err.Error())
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		logger.Errorf("unexpected error while getting rows affected: %s", err.Error())
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, apperrors.NewBadRequest("incorrect task parameters")
	}

	if err != nil {
		logger.Errorf("unexpected error while getting task by id: %s", err.Error())
		return nil, err
	}
	return taskToDelete, nil
}

func (s *storage) UpdateTaskGoal(userId, taskId, goalId int) (*task.Task, error) {

	query := `UPDATE tasks 
				SET 
					goal_id = ?,
					updated_at = CURRENT_TIMESTAMP
				WHERE id = ? 
				  AND user_id = (
				      SELECT g.user_id 
				      FROM goals g 
				      WHERE g.id = ? 
				        AND g.user_id = ?  
				        AND deleted_at IS NULL LIMIT 1)
				  AND deleted_at IS NULL`

	result, err := s.client.Exec(
		query,
		goalId,
		taskId,
		goalId,
		userId,
	)

	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("task has not updated")
	}

	taskGoalUpdated, err := s.GetTaskById(userId, taskId)
	if err != nil {
		return nil, err
	}

	return taskGoalUpdated, nil
}

func (s *storage) DeleteTaskGoal(userId, taskId int) (*task.Task, error) {

	query := `UPDATE tasks 
				SET 
					goal_id = NULL,
					updated_at = CURRENT_TIMESTAMP
				WHERE id = ? 
				  AND user_id = ?
				  AND goal_id IS NOT NULL
				  AND deleted_at IS NULL`

	result, err := s.client.Exec(
		query,
		taskId,
		userId,
	)

	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("task has not updated")
	}

	taskGoalUpdated, err := s.GetTaskById(userId, taskId)
	if err != nil {
		return nil, err
	}

	return taskGoalUpdated, nil
}
