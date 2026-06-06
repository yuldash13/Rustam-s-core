package domain

type User struct {
	ID          int64
	Name        string
	PhoneNumber string
	Mail        string
}

type UsersFilter struct {
	Name        string
	PhoneNumber string
	Mail        string
	Limit       int32
}

type HttpUser struct {
	Name string
	Role Role
}

type Role string

const (
	AdminRole  Role = "admin"
	BankirRole Role = "bank"
	UserRole   Role = "user"
)
