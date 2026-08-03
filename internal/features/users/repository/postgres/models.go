package users_postgres_repository

type UserModel struct {
	Id          int
	Version     int
	Username    string
	PhoneNumber *string
}
