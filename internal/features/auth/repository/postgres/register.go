package auth_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Dex564/golang-todoapp/internal/core/domain"
)

func (r *AuthRepository) Register(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO todoapp.users (username, login, email, password_hash)
	VALUES ($1, $2, $3, $4)
	RETURNING id, version, username, login, email, password_hash;`

	var userModel UserModel
	row := r.pool.QueryRow(ctx, query, user.Username, user.Login, user.Email, user.PasswordHash)
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.Username,
		&userModel.Login,
		&userModel.Email,
		&userModel.PasswordHash,
	)
	if err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.Username,
		userModel.Login,
		userModel.Email,
		userModel.PasswordHash,
	)

	return userDomain, nil
}
