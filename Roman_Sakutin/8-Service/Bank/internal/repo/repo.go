package repo

type Repo struct {
	Users     Users
	Accounts  Accounts
	Transfers Transfers
}

func NewRepo() *Repo {
	return &Repo{
		Users:     newUsers(),
		Accounts:  newAccounts(),
		Transfers: newTransfers(),
	}
}
