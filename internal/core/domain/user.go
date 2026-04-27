package domain

type User struct {
	ID      int
	Version int

	Username    string
	PhoneNumber *string
}
