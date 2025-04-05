package contract

import "calpal-core/entity"

type AuthRepository interface {
	SignUp(user entity.User) error
	SignIn() (entity.User, error)
}
