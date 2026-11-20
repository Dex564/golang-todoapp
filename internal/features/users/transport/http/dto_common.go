package users_transport_http

import "github.com/Dex564/golang-todoapp/internal/core/domain"

type UserDTOResponse struct {
	Id       int    `json:"id" example:"10"`
	Version  int    `json:"version" example:"3"`
	Username string `json:"username" example:"Ivan Ivanov"`
	Login    string `json:"login"`
	Email    string `json:"email"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		Id:       user.ID,
		Version:  user.Version,
		Username: user.Username,
		Login:    user.Login,
		Email:    user.Email,
	}
}

func usersDTOFromDomains(users []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))

	for i, user := range users {
		usersDTO[i] = userDTOFromDomain(user)
	}

	return usersDTO
}
