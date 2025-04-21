package auth

import (
	"calpal-core/entity"
	errhandle "calpal-core/pkg/err_handle"
	"calpal-core/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	repo repository.AuthRepository
}

func NewAuthHandler(r repository.AuthRepository) Handler {
	return Handler{repo: r}
}

func (a *Handler) SignUp(e echo.Context) error {

	var signUpUser entity.User

	// Decode json body
	err := e.Bind(&signUpUser)
	if err != nil {
		return e.JSON(http.StatusBadRequest, errhandle.NewError(errhandle.BadRequestError))
	}

	// Validate sign up entity
	if err = ValidateUser(signUpUser); err != nil {
		return e.JSON(http.StatusBadRequest, errhandle.NewError(errhandle.BadRequestError))
	}

	uid, err := a.repo.SignUp(signUpUser)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errhandle.NewError(errhandle.UnexpectedError))
	}

	return e.JSON(http.StatusOK, uid)
}

func (a *Handler) SignIn(e echo.Context) error {
	var signInUser entity.User

	err := e.Bind(&signInUser)
	if err != nil {
		return e.JSON(http.StatusBadRequest, errhandle.NewError(errhandle.BadRequestError))
	}

	if err = ValidateUser(signInUser); err != nil {
		return e.JSON(http.StatusBadRequest, errhandle.NewError(errhandle.BadRequestError))
	}

	return e.JSON(http.StatusOK, entity.User{})
}
