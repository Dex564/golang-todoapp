package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Dex564/golang-todoapp/internal/core/domain"
	core_errors "github.com/Dex564/golang-todoapp/internal/core/errors"
	core_postgres_pool "github.com/Dex564/golang-todoapp/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) GetUser(
	ctx context.Context,
	id int,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, username, phone_number
	FROM todoapp.users
	WHERE id=$1;
	`

	row := r.pool.QueryRow(ctx, query, id)

	var userModel UserModel
	err := row.Scan(
		&userModel.Id,
		&userModel.Version,
		&userModel.Username,
		&userModel.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with id='%d': %w", id, core_errors.ErrNotFound)
		}

		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.Id,
		userModel.Version,
		userModel.Username,
		userModel.PhoneNumber,
	)

	return userDomain, nil
}
