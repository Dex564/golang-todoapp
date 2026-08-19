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

	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)

	GetUser(
		ctx context.Context,
		id int,
	) (domain.User, error)

	DeleteUser(
		ctx context.Context,
		id int,
	) error

	PatchUser(
		ctx context.Context,
		id int,
		user domain.User,
	) (domain.User, error)
}

func NewUsersService(userRepository UserRepository) *UsersService {
	return &UsersService{
		usersRepository: userRepository,
	}
}
