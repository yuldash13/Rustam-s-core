package repo

import (
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/db"
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type Accounts interface {
	GetAccounts(ctx context.Context, IDUser int) ([]db.Account, error)
	GetAccountByCurrency(ctx context.Context, id int, currency string) (*db.Account, error)
	GetAccountByID(ctx context.Context, id int) (*db.Account, error)
	CreateAccount(ctx context.Context, a *db.Account) (int, error)
	UpdateAccount(ctx context.Context, id int, balance int) error
	DepAccount(ctx context.Context, dep db.DepAccount, id int) error
	DeleteAccount(ctx context.Context, id int) error
}

type AccountsRepo struct {
	db *pgxpool.Pool
}

func NewAccountRepo(db *pgxpool.Pool) *AccountsRepo {
	return &AccountsRepo{db: db}
}

func (AR *AccountsRepo) GetAccounts(ctx context.Context, IDUser int) ([]db.Account, error) {
	query := sq.Select(
		"id",
		"id_user",
		"balance",
		"currency",
	).From("accounts").OrderBy("created_at desc").Where(sq.Eq{"id_user": IDUser})
	query = query.PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := AR.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var items []db.Account
	for rows.Next() {
		var item db.Account
		if err = rows.Scan(&item.ID, &item.IDUser, &item.Balance, &item.Currency); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (AR *AccountsRepo) GetAccountByCurrency(ctx context.Context, id int, currency string) (*db.Account, error) {
	query := sq.Select(
		"id",
		"id_user",
		"balance",
		"currency",
	).From("accounts").
		OrderBy("created_at desc").
		Where(sq.And{sq.Eq{"id": id}, sq.Eq{"currency": currency}}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	row := AR.db.QueryRow(ctx, sql, args...)

	item := &db.Account{}
	err = row.Scan(&item.ID, &item.IDUser, &item.Balance, &item.Currency)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, db.NotFound
		}
		return nil, err
	}
	return item, nil
}

func (AR *AccountsRepo) GetAccountByID(ctx context.Context, id int) (*db.Account, error) {
	query := sq.Select(
		"id",
		"id_user",
		"balance",
		"currency",
	).From("accounts").
		OrderBy("created_at desc").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	row := AR.db.QueryRow(ctx, sql, args...)

	item := &db.Account{}
	err = row.Scan(&item.ID, &item.IDUser, &item.Balance, &item.Currency)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, db.NotFound
		}
		return nil, err
	}
	return item, nil
}

func (AR *AccountsRepo) CreateAccount(ctx context.Context, a *db.Account) (int, error) {
	createAt := time.Now()
	sql, args, err := sq.Insert("accounts").
		SetMap(map[string]interface{}{
			"id_user":    a.IDUser,
			"balance":    a.Balance,
			"currency":   a.Currency,
			"created_at": createAt,
		}).Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return 0, err
	}

	var row pgx.Row

	tx, ok := getTx(ctx)
	if ok {
		row = tx.QueryRow(ctx, sql, args...)
	} else {
		row = AR.db.QueryRow(ctx, sql, args...)
	}

	var id int
	if err = row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (AR *AccountsRepo) UpdateAccount(ctx context.Context, id int, balance int) error {
	update := sq.Update("accounts").
		Set("balance", sq.Expr("?", balance)).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := update.ToSql()
	if err != nil {
		return err
	}

	_, err = AR.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}

func (AR *AccountsRepo) DepAccount(ctx context.Context, dep db.DepAccount, id int) error {
	update := sq.Update("accounts").
		Set("balance", sq.Expr("balance + ?", dep.Dep)).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := update.ToSql()
	if err != nil {
		return err
	}

	_, err = AR.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}

func (AR *AccountsRepo) DeleteAccount(ctx context.Context, id int) error {
	del := sq.Delete("accounts").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := del.ToSql()
	if err != nil {
		return err
	}

	_, err = AR.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}
