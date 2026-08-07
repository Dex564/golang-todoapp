package users_transport_http

import "github.com/Dex564/golang-todoapp/internal/core/domain"

type UserDTOResponse struct {
	Id          int     `json:"id"`
	Version     int     `json:"version"`
	Username    string  `json:"username"`
	PhoneNumber *string `json:"phone_number"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		Id:          user.ID,
		Version:     user.Version,
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
	}
}

func usersDTOFromDomains(users []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))

	for i, user := range users {
		usersDTO[i] = userDTOFromDomain(user)
	}

	return usersDTO
}
