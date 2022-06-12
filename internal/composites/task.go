package composites

import (
	"doneclub-api/internal/adapters/api"
	task2 "doneclub-api/internal/adapters/api/task"
	task3 "doneclub-api/internal/adapters/db/task"
	"doneclub-api/internal/domain/task"
)

type TaskComposite struct {
	Storage task.Storage
	Service task2.Service
	Handler api.Handler
}

func NewTaskComposite(dbComposite *MySQLComposite) (*TaskComposite, error) {
	storage := task3.NewStorage(dbComposite.client)
	service := task.NewService(storage)
	handler := task2.NewHandler(service)

	return &TaskComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
