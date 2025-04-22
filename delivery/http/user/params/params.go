package params

import "calpal-core/entity"

type UsersRequest struct{}

type UsersResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
}

func ToUserResponse(user entity.User) UsersResponse {
	return UsersResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Username:  user.Username,
	}
}
