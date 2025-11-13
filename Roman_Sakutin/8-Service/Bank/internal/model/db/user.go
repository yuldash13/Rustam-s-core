package db

type User struct {
	ID          int
	Name        string
	PhoneNumber string
	Mail        string
}

type UsersFilter struct {
	Name        string
	PhoneNumber string
	Mail        string
	Limit       int
}
