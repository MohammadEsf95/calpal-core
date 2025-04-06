package contract

import "calpal-core/entity"

type AuthRepository interface {
	SignUp(user entity.User) (string, error)
	SignIn() (entity.User, error)
}
