package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/Dex564/golang-todoapp/internal/core/errors"
)

type User struct {
	ID      int
	Version int

	Username string
	Login    string
	Email    string

	PasswordHash string
}

func NewUser(id int, version int, username string, login string, email string, passwordHash string) User {
	return User{
		ID:           id,
		Version:      version,
		Username:     username,
		Login:        login,
		Email:        email,
		PasswordHash: passwordHash,
	}
}

func NewUserUninitialized(username string, login string, email string, passwordHash string) User {
	return NewUser(UninitializedID, UninitializedVersion, username, login, email, passwordHash)
}

func (u *User) Validate() error {
	usernameLength := len([]rune(u.Username))
	if usernameLength < 3 || usernameLength > 100 {
		return fmt.Errorf("invalid `Username` length: %d: %w", usernameLength, core_errors.ErrInvalidArgument)
	}

	loginLength := len([]rune(u.Login))
	if loginLength < 3 || loginLength > 30 {
		return fmt.Errorf("invalid `Login` length: %d: %w", loginLength, core_errors.ErrInvalidArgument)
	}

	emailLength := len([]rune(u.Email))
	if emailLength < 8 || emailLength > 254 {
		return fmt.Errorf("invalid `Email` length: %d: %w", emailLength, core_errors.ErrInvalidArgument)
	}
	re := regexp.MustCompile(`[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)
	if !re.MatchString(u.Email) {
		return fmt.Errorf("invalid `Email` format: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

type UserPatch struct {
	Username Nullable[string]
	Login    Nullable[string]
	Email    Nullable[string]
}

func NewUserPatch(
	username Nullable[string],
	login Nullable[string],
	email Nullable[string],
) UserPatch {
	return UserPatch{
		Username: username,
		Login:    login,
		Email:    email,
	}
}

func (p *UserPatch) Validate() error {
	if p.Username.Set && p.Username.Value == nil {
		return fmt.Errorf("`Username` can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}

	if p.Login.Set && p.Login.Value == nil {
		return fmt.Errorf("`Login` can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}

	if p.Email.Set && p.Email.Value == nil {
		return fmt.Errorf("`Email` can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate user patch: %w", err)
	}

	tmp := *u

	if patch.Username.Set {
		tmp.Username = *patch.Username.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched user: %w", err)
	}

	*u = tmp
	return nil
}
