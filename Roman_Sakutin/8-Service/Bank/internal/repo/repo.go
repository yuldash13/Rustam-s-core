package repo

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repo struct {
	Users       Users
	Accounts    Accounts
	Transfers   Transfers
	Transaction Transaction
}

func NewRepo(pool *pgxpool.Pool) *Repo {
	return &Repo{
		Users:       NewUserRepo(pool),
		Accounts:    NewAccountRepo(pool),
		Transfers:   NewTransferRepo(pool),
		Transaction: NewTransactionRepo(pool),
	}
}
