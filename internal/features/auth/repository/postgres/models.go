package auth_postgres_repository

type UserModel struct {
	ID      int
	Version int

	Username string
	Login    string
	Email    string

	PasswordHash string
}
