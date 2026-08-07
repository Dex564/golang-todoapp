package users_postgres_repository

import "github.com/Dex564/golang-todoapp/internal/core/domain"

type UserModel struct {
	Id          int
	Version     int
	Username    string
	PhoneNumber *string
}

func UserDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))

	for i, user := range users {
		userDomains[i] = domain.NewUser(
			user.Id,
			user.Version,
			user.Username,
			user.PhoneNumber,
		)
	}

	return userDomains
}
