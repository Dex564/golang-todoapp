package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/Dex564/golang-todoapp/internal/core/errors"
)

type User struct {
	ID      int
	Version int

	Username    string
	PhoneNumber *string
}

func NewUser(id int, version int, username string, phoneNumber *string) User {
	return User{
		ID:          id,
		Version:     version,
		Username:    username,
		PhoneNumber: phoneNumber,
	}
}

func NewUserUninitialized(username string, phoneNumber *string) User {
	return NewUser(UninitializedID, UninitializedVersion, username, phoneNumber)
}

func (u *User) Validate() error {
	usernameLength := len([]rune(u.Username))
	if usernameLength < 3 || usernameLength > 100 {
		return fmt.Errorf("invalid `Username` length: %d: %w", usernameLength, core_errors.ErrInvalidArgument)
	}

	if u.PhoneNumber != nil {
		phoneNumberLength := len([]rune(*u.PhoneNumber))
		if phoneNumberLength < 10 || phoneNumberLength > 15 {
			return fmt.Errorf("invalid `PhoneNumber` length: %d: %w", phoneNumberLength, core_errors.ErrInvalidArgument)
		}
		re := regexp.MustCompile(`^\+[0-9]+$`)
		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf("invalid `PhoneNumber` format: %w", core_errors.ErrInvalidArgument)
		}
	}

	return nil
}
