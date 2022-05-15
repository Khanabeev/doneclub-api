package composites

import (
	"doneclub-api/internal/adapters/api"
	user3 "doneclub-api/internal/adapters/api/user"
	user2 "doneclub-api/internal/adapters/db/user"
	"doneclub-api/internal/domain/user"
)

type AuthorComposite struct {
	Storage user.Storage
	Service user.Service
	Handler api.Handler
}

func NewUserComposite(dbComposite *MySQLComposite) (*AuthorComposite, error) {
	storage := user2.NewStorage(dbComposite.client)
	service := user.NewService(storage)
	handler := user3.NewHandler(service)

	return &AuthorComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
