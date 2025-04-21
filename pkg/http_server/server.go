package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Router *echo.Echo
	Config Config
}

type Config struct {
	Address string `koanf:"address"`
	Port    int    `koanf:"port"`
}

func New(cfg Config) Server {
	e := echo.New()

	e.Use(middleware.Logger())

	return Server{Router: e, Config: cfg}
}

func (s Server) Start() error {
	return s.Router.Start(fmt.Sprintf(":%d", s.Config.Port))
}
