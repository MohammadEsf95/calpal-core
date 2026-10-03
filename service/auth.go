package service

import (
	"calpal-core/delivery/http/auth/params"
	"calpal-core/entity"
	"calpal-core/repository"
	"github.com/google/uuid"
	"time"
)

type AuthService interface {
	SignUp(user params.SignUpRequest) (string, error)
	SignIn(req params.SignInRequest) (params.SignInResponse, error)
}

type authService struct {
	authRepo repository.AuthRepository
}

func NewAuthService(authRepo repository.AuthRepository) AuthService {
	return authService{authRepo: authRepo}
}

func (a authService) SignUp(user params.SignUpRequest) (string, error) {
	u := entity.User{
		ID:             uuid.New().String(),
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Username:       user.Username,
		Password:       user.Password,
		Email:          user.Email,
		TargetCalories: 0,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	return a.authRepo.SignUp(u)
}

func (a authService) SignIn(request params.SignInRequest) (params.SignInResponse, error) {
	a.authRepo.SignIn(entity.User{
		ID:             "",
		FirstName:      "",
		LastName:       "",
		Username:       "",
		Password:       "",
		Email:          "",
		TargetCalories: 0,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	})

	return params.SignInResponse{}, nil
}
