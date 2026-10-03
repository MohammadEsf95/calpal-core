package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Server) healthCheck(e echo.Context) error {
	return e.JSON(http.StatusOK, echo.Map{
		"message": "active",
	})
}
