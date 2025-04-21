package http

import (
	"calpal-core/delivery/http/auth"
	"calpal-core/delivery/http/user"
	httpserver "calpal-core/pkg/http_server"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	HTTPServer  httpserver.Server
	AuthHandler auth.Handler
	UserHandler user.Handler
}

func New(server httpserver.Server, authHandler auth.Handler, userHandler user.Handler) Server {
	return Server{
		HTTPServer:  server,
		AuthHandler: authHandler,
		UserHandler: userHandler,
	}
}

func (s Server) Serve() error {
	s.RegisterRoutes()
	if err := s.HTTPServer.Start(); err != nil {
		return err
	}

	return nil
}

func (s Server) RegisterRoutes() {

	s.HTTPServer.Router.Use(middleware.Recover())

	v1 := s.HTTPServer.Router.Group("/v1")

	v1.GET("/health-check", s.healthCheck)
}
