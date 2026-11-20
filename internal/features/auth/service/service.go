package auth_service

import (
	"context"

	"github.com/Dex564/golang-todoapp/internal/core/domain"
	"github.com/Dex564/golang-todoapp/internal/core/security"
)

type AuthService struct {
	authRepository AuthRepository
	hasher         security.PasswordHasher
}

type AuthRepository interface {
	Register(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
}

func NewAuthService(authRepository AuthRepository) *AuthService {
	return &AuthService{authRepository: authRepository}
}
