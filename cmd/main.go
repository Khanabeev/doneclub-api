package main

import (
	"doneclub-api/internal/composites"
	"doneclub-api/pkg/config"
	"doneclub-api/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	logger.Info("load configuration")
	err := config.Load(".env")
	if err != nil {
		logger.Fatal(err)
		panic(err)
	}

	logger.Info("start Data Base composite")
	dbComposite, err := composites.NewMySQLComposite()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("start User composite")
	userComposite, err := composites.NewUserComposite(dbComposite)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("register user handler")
	userComposite.Handler.Register(router)

	logger.Info("start application...")
	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("start application")

	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Info("server is listening port 1234")
	log.Fatal(server.Serve(listener))
}
