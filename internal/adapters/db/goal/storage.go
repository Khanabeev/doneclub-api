package goal

import (
	"doneclub-api/internal/domain/goal"
	"doneclub-api/internal/domain/user"
	"doneclub-api/pkg/apperrors"
	"doneclub-api/pkg/logging"
	"github.com/jmoiron/sqlx"
)

type storage struct {
	client *sqlx.DB
}

func NewStorage(client *sqlx.DB) goal.Storage {
	return &storage{
		client: client,
	}
}

func (s *storage) GetAllGoalsByUser(user *user.User, limit int, offset int) ([]*goal.Goal, error) {
	return nil, nil
}
func (s *storage) GetGoalById(userId, goalId int) (*goal.Goal, error) {
	return nil, nil
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

func (s *storage) UpdateGoal(g *goal.Goal, goalId string) (*goal.Goal, error) {
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
		logger.Errorf("unexpected error while getting last inserted goal id: %s", err.Error())
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, apperrors.NewBadRequest("incorrect goal parameters")
	}

	return g, nil
}
