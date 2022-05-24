package composites

import (
	"doneclub-api/internal/adapters/api"
	user3 "doneclub-api/internal/adapters/api/user"
	user2 "doneclub-api/internal/adapters/db/user"
	"doneclub-api/internal/domain/user"
)

type GoalComposite struct {
	Storage user.Storage
	Service user.Service
	Handler api.Handler
}

func NewGoalComposite(dbComposite *MySQLComposite) (*GoalComposite, error) {
	storage := user2.NewStorage(dbComposite.client)
	service := user.NewService(storage)
	handler := user3.NewHandler(service)

	return &GoalComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
