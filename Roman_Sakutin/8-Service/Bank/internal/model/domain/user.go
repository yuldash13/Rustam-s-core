package domain

type GetUsersResponse struct {
	Users []User
}

type GetUsersByID struct {
	User     *User
	Accounts GetAccountsResponse
}

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Mail        string `json:"mail"`
}

type UsersFilterRequest struct {
	Name        string `form:"name"`
	PhoneNumber string `form:"phone_number"`
	Mail        string `form:"mail"`
	Limit       int    `form:"limit"`
}
