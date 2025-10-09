package repo

import (
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/db"
	"context"
	"errors"
	"math/rand"
	"sync"
)

type Transfers interface {
	GetTransferByID(ctx context.Context, id int, limit int) ([]db.Transfer, error)
	MakeTransfer(ctx context.Context, t *db.Transfer) (int, error)
	CancelTransfer(ctx context.Context, id int) (*db.Transfer, error)
}

type transfers struct {
	data map[int]*db.Transfer
	m    sync.RWMutex
}

func newTransfers() *transfers {
	return &transfers{
		data: make(map[int]*db.Transfer, 1000000),
		m:    sync.RWMutex{},
	}
}

func (rT *transfers) GetTransferByID(ctx context.Context, id int, limit int) ([]db.Transfer, error) {
	rT.m.RLock()
	defer rT.m.RUnlock()

	var items = make([]db.Transfer, 0, 5)

	for _, r := range rT.data {
		if limit == 0 {
			break
		}
		if r.IDFrom == id || r.IDTo == id {
			items = append(items, *r)
			limit--
		}
	}
	return items, nil
}

func (rT *transfers) MakeTransfer(ctx context.Context, t *db.Transfer) (int, error) {
	id := rand.Intn(1000000)

	rT.m.RLock()
	defer rT.m.RUnlock()

	if _, ok := rT.data[id]; ok {
		return 0, errors.New("failed to create")
	}
	t.ID = id
	t.OperationState =db.Success
	rT.data[id] = t
	return id, nil
}

func (rT *transfers) CancelTransfer(ctx context.Context, id int) (*db.Transfer, error) {
	rT.m.RLock()
	defer rT.m.RUnlock()

	var item *db.Transfer

	for _, r := range rT.data {
		if r.ID == id {
			r.OperationState = db.Canceled
			item = r
		}
	}
	return item, nil
}
