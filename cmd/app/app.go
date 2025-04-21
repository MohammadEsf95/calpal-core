package main

import (
	"calpal-core/delivery/http"
	"calpal-core/delivery/http/auth"
	"calpal-core/delivery/http/user"
	httpserver "calpal-core/pkg/http_server"
	"calpal-core/repository"
	"database/sql"
)

type Application struct {
	server http.Server
}

func Setup(db *sql.DB, cfg Config) Application {
	authRepo := repository.NewAuthRepository(db)
	authHandler := auth.NewAuthHandler(authRepo)

	userRepo := repository.NewUserRepository(db)
	userHandler := user.New(userRepo)

	server := httpserver.New(cfg.HttpServer)
	httpServer := http.New(server, authHandler, userHandler)

	return Application{
		httpServer,
	}
}

func (a Application) Start() {
	a.server.Serve()
}
