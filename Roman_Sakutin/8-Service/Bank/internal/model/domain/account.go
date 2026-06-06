package domain

type Account struct {
	ID       int64
	IDUser   int64
	Balance  int64
	Currency string
}

type DepAccount struct {
	Dep int32
}
