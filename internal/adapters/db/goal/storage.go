package goal

import (
	"doneclub-api/internal/domain/goal"
	"doneclub-api/internal/domain/user"
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
