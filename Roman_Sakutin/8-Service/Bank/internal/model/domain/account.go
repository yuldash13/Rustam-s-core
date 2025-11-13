package domain

type GetAccountsResponse struct {
	Accounts []Account
}

type GetAccountsByID struct {
	Account *Account
}

type DepAccount struct {
	Dep int `json:"dep"`
}

type Account struct {
	ID       int    `json:"id"`
	IDUser   int    `json:"id_user"`
	Balance  int    `json:"balance"`
	Currency string `json:"currency"`
}
