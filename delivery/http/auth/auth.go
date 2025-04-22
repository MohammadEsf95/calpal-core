package auth

import (
	"calpal-core/delivery/http/auth/params"
	errhandle "calpal-core/pkg/err_handle"
	"calpal-core/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	service   service.AuthService
	validator Validator
}

func NewAuthHandler(r service.AuthService, v Validator) Handler {
	return Handler{service: r, validator: v}
}

func (a *Handler) SignUp(e echo.Context) error {

	var signUpUser params.SignUpRequest

	// Decode json body
	err := e.Bind(&signUpUser)
	if err != nil {
		return e.JSON(http.StatusBadRequest, errhandle.NewError(errhandle.BadRequestError))
	}

	//Validate sign up entity
	if err = a.validator.ValidateSignUp(signUpUser); err != nil {
		return e.JSON(http.StatusBadRequest, errhandle.NewError(errhandle.BadRequestError))
	}

	uid, err := a.service.SignUp(signUpUser)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errhandle.NewError(errhandle.UnexpectedError, map[string]interface{}{
			"error": err.Error(),
		}))
	}

	return e.JSON(http.StatusOK, uid)
}

func (a *Handler) SignIn(e echo.Context) error {
	var signInUser params.SignInRequest

	err := e.Bind(&signInUser)
	if err != nil {
		return e.JSON(http.StatusBadRequest, errhandle.NewError(errhandle.BadRequestError))
	}

	if err = a.validator.ValidateSignIn(signInUser); err != nil {
		return e.JSON(http.StatusBadRequest, errhandle.NewError(errhandle.BadRequestError))
	}

	in, err := a.service.SignIn(params.SignInRequest{})
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errhandle.NewError(errhandle.UnexpectedError))
	}

	return e.JSON(http.StatusOK, in)
}
