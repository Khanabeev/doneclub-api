package composites

import (
	"doneclub-api/pkg/client/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLComposite struct {
	client *sqlx.DB
}

func NewMySQLComposite() (*MySQLComposite, error) {
	client, err := mysql.NewClient()
	if err != nil {
		return nil, err
	}

	return &MySQLComposite{
		client: client,
	}, nil

}
