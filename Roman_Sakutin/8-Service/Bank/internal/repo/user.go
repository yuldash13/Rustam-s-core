package repo

import (
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/db"
	"context"
	"errors"
	"math/rand"
	"strings"
	"sync"
)

type Users interface {
	GetUsers(ctx context.Context, filters *db.UsersFilter) ([]db.User, error)
	GetUserByID(ctx context.Context, id int) (*db.User, error)
	CreateUser(ctx context.Context, u *db.User) (int, error)
	UpdateUser(ctx context.Context, id int, u *db.User) error
}

type users struct {
	data map[int]*db.User
	m    sync.RWMutex
}

func newUsers() *users {
	return &users{
		data: make(map[int]*db.User, 1000000),
		m:    sync.RWMutex{},
	}
}

func (rU *users) GetUsers(ctx context.Context, filters *db.UsersFilter) ([]db.User, error) {
	rU.m.RLock()
	defer rU.m.RUnlock()

	var items = make([]db.User, 0, len(rU.data))

	for _, item := range rU.data {
		if filtersSearch(item, filters) == true {
			items = append(items, *item)
		}
	}
	return items, nil
}

func filtersSearch(user *db.User, filters *db.UsersFilter) bool {
	if filters.Name != "" {
		if !strings.Contains(strings.ToLower(user.Name), strings.ToLower(filters.Name)) {
			return false
		}
	}
	if filters.PhoneNumber != "" {
		if !strings.Contains(strings.ToLower(user.PhoneNumber), strings.ToLower(filters.PhoneNumber)) {
			return false
		}
	}
	if filters.Mail != "" {
		if !strings.Contains(strings.ToLower(user.Mail), strings.ToLower(filters.Mail)) {
			return false
		}
	}
	return true
}

func (rU *users) GetUserByID(ctx context.Context, id int) (*db.User, error) {
	rU.m.RLock()
	defer rU.m.RUnlock()

	for _, item := range rU.data {
		if item.ID == id {
			return item, nil
		}
	}
	return nil, db.NotFound
}

func (rU *users) CreateUser(ctx context.Context, u *db.User) (int, error) {
	id := rand.Intn(1000000)

	rU.m.Lock()
	defer rU.m.Unlock()

	if _, ok := rU.data[id]; ok {
		return 0, errors.New("failed to create")
	}
	u.ID = id
	rU.data[id] = u
	return id, nil
}

func (rU *users) UpdateUser(ctx context.Context, id int, u *db.User) error {
	if rU.data[id] == nil {
		return db.NotFound
	}
	rU.data[id] = u
	return nil
}
