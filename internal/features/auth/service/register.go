package auth_service

import (
	"context"
	"fmt"

	"github.com/Dex564/golang-todoapp/internal/core/domain"
)

func (s *AuthService) Register(
	ctx context.Context,
	username string,
	login string,
	email string,
	password string,
) (domain.User, error) {

	passwordHash, err := s.hasher.Hash(password)
	if err != nil {
		return domain.User{}, fmt.Errorf("hash password: %w", err)
	}

	userDomain := domain.NewUserUninitialized(username, login, email, passwordHash)
	if err := userDomain.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user: %w", err)
	}

	return s.authRepository.Register(ctx, userDomain)
}
