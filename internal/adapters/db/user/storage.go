package user

import (
	"database/sql"
	"doneclub-api/internal/domain/user"
	"doneclub-api/pkg/logging"
	"github.com/jmoiron/sqlx"
)

type storage struct {
	client *sqlx.DB
}

func NewStorage(client *sqlx.DB) user.Storage {
	return &storage{
		client: client,
	}
}

func (s *storage) CreateUser(user *user.User) (*user.User, error) {
	logger := logging.GetLogger()
	query := "INSERT INTO doneclub.users (email, password, status, created_at, updated_at) VALUE (?, ?, ?, ?, ?)"
	result, err := s.client.Exec(query, user.Email, user.Password, user.Status, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	user.ID = int(id)

	return user, nil
}
func (s *storage) GetUserById(userId int) (*user.User, error) {
	return nil, nil
}

func (s *storage) GetUserByEmail(userEmail string) (*user.User, error) {
	var u user.User

	query := "SELECT email, status FROM doneclub.users WHERE email = ?"
	err := s.client.Get(&u, query, userEmail)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &u, nil

}
