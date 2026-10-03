package main

import (
	"calpal-core/delivery/http"
	"calpal-core/delivery/http/auth"
	"calpal-core/delivery/http/user"
	httpserver "calpal-core/pkg/http_server"
	"calpal-core/repository"
	"calpal-core/service"
	"database/sql"
)

type Application struct {
	server http.Server
}

func Setup(db *sql.DB, cfg Config) Application {
	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(authService, auth.NewValidator())

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := user.New(userService)

	server := httpserver.New(cfg.HttpServer)
	httpServer := http.New(server, authHandler, userHandler)

	return Application{
		httpServer,
	}
}

func (a Application) Start() {
	a.server.Serve()
}
