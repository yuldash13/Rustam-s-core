package repo

import (
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/db"
	"context"
	"errors"
	"math/rand"
	"sync"
)

type Accounts interface {
	GetAccounts(ctx context.Context, id int) ([]db.Account, error)
	GetAccountByCurrency(ctx context.Context, id int, currency string) (*db.Account, error)
	GetAccountByID(ctx context.Context, id int) (*db.Account, error)
	CreateAccount(ctx context.Context, a *db.Account) (int, error)
	DepAccount(ctx context.Context, dep int, id int) error
	DeleteAccount(ctx context.Context, id int) error
}

type accounts struct {
	data map[int]*db.Account
	m    sync.RWMutex
}

func newAccounts() *accounts {
	return &accounts{
		data: make(map[int]*db.Account, 1000000),
		m:    sync.RWMutex{},
	}
}

func (rA *accounts) GetAccounts(ctx context.Context, id int) ([]db.Account, error) {
	rA.m.RLock()
	defer rA.m.RUnlock()

	var items = make([]db.Account, 0, len(rA.data))

	for _, item := range rA.data {
		if item.IDUser == id {
			items = append(items, *item)
		}
	}
	return items, nil
}

func (rA *accounts) GetAccountByCurrency(ctx context.Context, id int, currency string) (*db.Account, error) {
	rA.m.RLock()
	defer rA.m.RUnlock()

	var item *db.Account

	for _, r := range rA.data {
		if r.IDUser == id {
			if r.Currency == currency {
				item = r
				break
			}
		}
	}
	return item, nil
}

func (rA *accounts) GetAccountByID(ctx context.Context, id int) (*db.Account, error) {
	rA.m.RLock()
	defer rA.m.RUnlock()

	for _, item := range rA.data {
		if item.ID == id {
			return item, nil
		}
	}
	return nil, db.NotFound
}

func (rA *accounts) CreateAccount(ctx context.Context, a *db.Account) (int, error) {
	items, _ := rA.GetAccounts(ctx, a.IDUser)
	if ok, err := checkAcc(items, a.Currency); ok == true {
		return 0, err
	}

	id := rand.Intn(1000000)

	rA.m.Lock()
	defer rA.m.Unlock()

	if _, ok := rA.data[id]; ok {
		return 0, errors.New("failed to create")
	}
	a.ID = id
	rA.data[id] = a
	return id, nil
}

func checkAcc(items []db.Account, currency string) (bool, error) {
	if currency == "ru" || currency == "usd" || currency == "eu" {
		for _, r := range items {
			if r.Currency == currency {
				return true, db.Exist
			}
		}
		return false, nil
	}
	return true, db.Exist
}

func (rA *accounts) DepAccount(ctx context.Context, dep int, id int) error {
	if rA.data[id] == nil {
		return db.NotFound
	}
	rA.data[id].Balance += dep
	return nil
}

func (rA *accounts) DeleteAccount(ctx context.Context, id int) error {
	if rA.data[id] == nil {
		return db.NotFound
	}
	delete(rA.data, id)
	return nil
}
