package service

import (
	"calpal-core/delivery/http/user/params"
	"calpal-core/repository"
)

type UserService interface {
	Users() ([]params.UsersResponse, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return userService{repo: repo}
}

func (u userService) Users() ([]params.UsersResponse, error) {
	users, err := u.repo.Users()
	if err != nil {
		return nil, err
	}

	var res []params.UsersResponse
	for _, user := range users {
		res = append(res, params.ToUserResponse(user))
	}

	return res, nil
}
