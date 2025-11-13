package db

type Account struct {
	ID       int
	IDUser   int
	Balance  int
	Currency string
}

type DepAccount struct {
	Dep int
}
