package composites

import (
	"doneclub-api/internal/adapters/api"
	goal3 "doneclub-api/internal/adapters/api/goal"
	"doneclub-api/internal/adapters/db/goal"
	goal2 "doneclub-api/internal/domain/goal"
)

type GoalComposite struct {
	Storage goal2.Storage
	Service goal3.Service
	Handler api.Handler
}

func NewGoalComposite(dbComposite *MySQLComposite) (*GoalComposite, error) {
	storage := goal.NewStorage(dbComposite.client)
	service := goal2.NewService(storage)
	handler := goal3.NewHandler(service)

	return &GoalComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
