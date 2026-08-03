package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Dex564/golang-todoapp/internal/core/domain"
)

func (r *UsersRepository) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO todoapp.users (username, phone_number)
	VALUES ($1, $2)
	RETURNING id, version, username, phone_number;
	`

	row := r.pool.QueryRow(ctx, query, user.Username, user.PhoneNumber)
	var userModel UserModel
	err := row.Scan(
		&userModel.Id,
		&userModel.Version,
		&userModel.Username,
		&userModel.PhoneNumber,
	)
	if err != nil {
		return domain.User{}, fmt.Errorf("scan error %w", err)
	}

	userDomain := domain.NewUser(
		userModel.Id,
		userModel.Version,
		userModel.Username,
		userModel.PhoneNumber,
	)

	return userDomain, nil
}
