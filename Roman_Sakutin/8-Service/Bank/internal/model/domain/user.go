package domain

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Mail        string `json:"mail"`
}

type UsersFilter struct {
	Name        string `form:"name"`
	PhoneNumber string `form:"phone_number"`
	Mail        string `form:"mail"`
	Limit       int    `form:"limit"`
}
