package auth_transport_http

import "github.com/Dex564/golang-todoapp/internal/core/domain"

type UserDTOResponse struct {
	Id       int    `json:"id" example:"10"`
	Username string `json:"username" example:"Ivan Ivanov"`
	Login    string `json:"login"`
	Email    string `json:"email"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		Id:       user.ID,
		Username: user.Username,
		Login:    user.Login,
		Email:    user.Email,
	}
}
