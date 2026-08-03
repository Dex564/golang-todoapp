package users_service

import (
	"context"

	"github.com/Dex564/golang-todoapp/internal/core/domain"
)

type UsersService struct {
	usersRepository UserRepository
}

type UserRepository interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
}

func NewUsersService(userRepository UserRepository) *UsersService {
	return &UsersService{
		usersRepository: userRepository,
	}
}
