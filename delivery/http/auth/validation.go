package auth

import (
	"calpal-core/delivery/http/auth/params"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

var (
	ErrEmailRequired       = "email is required"
	ErrWrongEmailFormat    = "wrong email format"
	ErrUsernameRequired    = "username is required"
	ErrUsernameWrongLength = "username is too short or too long"
)

type Validator struct{}

func NewValidator() Validator { return Validator{} }

func (v Validator) ValidateSignUp(req params.SignUpRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.Email,
			validation.Required.Error(ErrEmailRequired),
			is.Email.Error(ErrWrongEmailFormat)),
		validation.Field(
			&req.Username,
			validation.Required.Error(ErrUsernameRequired),
			validation.Length(2, 20).Error(ErrUsernameWrongLength),
		),
		validation.Field(
			&req.FirstName,
			validation.Required,
			validation.Length(2, 20),
		),
		validation.Field(
			&req.LastName,
			validation.Required,
			validation.Length(2, 50),
		),
		validation.Field(
			&req.Password,
			validation.Length(8, 20),
		),
	)
}

func (v Validator) ValidateSignIn(req params.SignInRequest) error {

	return nil
}
