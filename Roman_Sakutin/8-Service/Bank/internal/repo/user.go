package repo

import (
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/db"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type Users interface {
	GetUsers(ctx context.Context, filters *db.UsersFilter) ([]db.User, error)
	GetUserByID(ctx context.Context, id int) (*db.User, []db.Account, error)
	CreateUser(ctx context.Context, u *db.User) (int, error)
	UpdateUser(ctx context.Context, u *db.User) error
}

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{db: db}
}

func (UR *UserRepo) GetUsers(ctx context.Context, filters *db.UsersFilter) ([]db.User, error) {
	query := `SELECT id, name, phone_number, mail FROM users WHERE 1=1`
	var args []interface{}

	nextArg := func(v interface{}) string {
		args = append(args, v)
		return fmt.Sprintf("$%d", len(args))
	}

	if filters != nil {
		if len(filters.Name) != 0 {
			query += " AND name ILIKE " + nextArg(filters.Name)
		}
		if len(filters.PhoneNumber) != 0 {
			query += " AND phone_number ILIKE " + nextArg(filters.PhoneNumber+"%")
		}
		if len(filters.Mail) != 0 {
			query += " AND mail ILIKE " + nextArg("%"+filters.Mail+"%")
		}
	}

	query += " ORDER BY created_at DESC"

	if filters != nil && filters.Limit != 0 {
		query += " LIMIT " + nextArg(filters.Limit)
	}

	rows, err := UR.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []db.User
	for rows.Next() {
		var item db.User
		if err = rows.Scan(&item.ID, &item.Name, &item.PhoneNumber, &item.Mail); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (UR *UserRepo) GetUserByID(ctx context.Context, id int) (*db.User, []db.Account, error) {
	query := `
		SELECT id, name, phone_number, mail
		FROM users
		WHERE id = $1
	`

	row := UR.db.QueryRow(ctx, query, id)

	item := &db.User{}
	err := row.Scan(&item.ID, &item.Name, &item.PhoneNumber, &item.Mail)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil, db.NotFound
		}
		return nil, nil, err
	}
	return item, nil, nil
}

func (UR *UserRepo) CreateUser(ctx context.Context, u *db.User) (int, error) {
	query := `
		INSERT INTO users (name, phone_number, mail, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	createAt := time.Now()
	args := []interface{}{u.Name, u.PhoneNumber, u.Mail, createAt}

	var row pgx.Row

	tx, ok := getTx(ctx)
	if ok {
		row = tx.QueryRow(ctx, query, args...)
	} else {
		row = UR.db.QueryRow(ctx, query, args...)
	}

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (UR *UserRepo) UpdateUser(ctx context.Context, u *db.User) error {
	update := `
		UPDATE users 
		SET name = $1, phone_number = $2, mail = $3 
		WHERE id = $4
	`

	args := []interface{}{u.Name, u.PhoneNumber, u.Mail, u.ID}

	res, err := UR.db.Exec(ctx, update, args...)

	if err != nil {
		return err
	}
	if n := res.RowsAffected(); n == 0 {
		return sql.ErrNoRows
	}
	return nil
}
