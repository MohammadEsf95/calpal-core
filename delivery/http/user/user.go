package user

import (
	"calpal-core/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	service service.UserService
}

func New(srv service.UserService) Handler {
	return Handler{service: srv}
}

func (h *Handler) Users(e echo.Context) error {

	users, err := h.service.Users()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, users)
}
