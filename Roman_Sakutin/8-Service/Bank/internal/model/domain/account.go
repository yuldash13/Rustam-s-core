package domain

type Account struct {
	ID       int    `json:"id"`
	IDUser   int    `json:"id_user"`
	Balance  int    `json:"balance"`
	Date     string `json:"date"`
	Currency string `json:"currency"`
}
