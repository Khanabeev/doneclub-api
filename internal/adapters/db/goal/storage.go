package goal

import (
	"doneclub-api/internal/domain/goal"
	"doneclub-api/pkg/apperrors"
	"doneclub-api/pkg/logging"
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

type storage struct {
	client *sqlx.DB
}

func NewStorage(client *sqlx.DB) goal.Storage {
	return &storage{
		client: client,
	}
}

func (s *storage) GetAllGoalsByUserIdAndStatus(userId, status int) ([]*goal.Goal, error) {
	query := `SELECT * 
				FROM goals
				WHERE user_id = ?
				  AND status = ?
				  AND deleted_at IS NULL`
	var goals []*goal.Goal
	err := s.client.Select(&goals, query, userId, status)
	if err != nil {
		return nil, err
	}

	return goals, nil
}

func (s *storage) GetAllGoalsByUserId(userId int) ([]*goal.Goal, error) {
	query := `SELECT * 
				FROM goals
				WHERE user_id = ?
				  AND deleted_at IS NULL`
	var goals []*goal.Goal
	err := s.client.Select(&goals, query, userId)
	if err != nil {
		return nil, err
	}

	return goals, nil
}

func (s *storage) GetGoalById(userId int, goalId int) (*goal.Goal, error) {
	query := `SELECT *
				FROM goals
				WHERE id = ? AND user_id = ? AND deleted_at IS NULL`
	var g goal.Goal
	err := s.client.Get(&g, query, goalId, userId)

	if err != nil {
		return nil, err
	}

	return &g, nil
}
func (s *storage) CreateGoal(g *goal.Goal) (*goal.Goal, error) {
	logger := logging.GetLogger()
	query := "INSERT INTO goals (user_id, status, parent_id, title, description,start_date, end_date) VALUES(?, ?, ?, ?, ?, ?, ?)"
	result, err := s.client.Exec(query, g.UserID, g.Status, g.ParentID, g.Title, g.Description, g.StartDate, g.EndDate)
	if err != nil {
		logger.Errorf("unexpected error during insert new goal: %s", err.Error())
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Errorf("unexpected error while getting last inserted goal id: %s", err.Error())
		return nil, err
	}

	g.ID = int(id)
	return g, nil
}

func (s *storage) UpdateGoal(g *goal.Goal, goalId int) (*goal.Goal, error) {
	logger := logging.GetLogger()
	query := `UPDATE goals 
				SET 
				    status = ?, 
				    parent_id = ?, 
				    title = ?, 
				    description = ?, 
				    start_date = ?, 
				    end_date = ?, 
				    updated_at = ?
				WHERE user_id = ? 
				  AND id = ? 
				  AND deleted_at IS NULL`

	result, err := s.client.Exec(
		query,
		g.Status,
		g.ParentID,
		g.Title,
		g.Description,
		g.StartDate,
		g.EndDate,
		g.UpdatedAt,
		g.UserID,
		goalId,
	)
	if err != nil {
		logger.Errorf("unexpected error during insert new goal: %s", err.Error())
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Errorf("unexpected error while getting rows affected: %s", err.Error())
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, apperrors.NewBadRequest("incorrect goal parameters")
	}

	return g, nil
}

func (s *storage) DeleteGoalById(userId int, goalId int) error {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	query := `UPDATE goals 
				SET deleted_at = ?
				WHERE id = ? AND user_id = ? AND deleted_at IS NULL`

	result, err := s.client.Exec(query, currentTime, goalId, userId)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}
