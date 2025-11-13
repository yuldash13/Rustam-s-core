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

type Transfers interface {
	GetTransfers(ctx context.Context, id int, limit int) ([]db.Transfer, error)
	GetTransferByID(ctx context.Context, id int) (*db.Transfer, error)
	MakeTransfer(ctx context.Context, t *db.Transfer) (int, error)
	CancelTransfer(ctx context.Context, id int) error
}

type TransfersRepo struct {
	db *pgxpool.Pool
}

func NewTransferRepo(db *pgxpool.Pool) *TransfersRepo {
	return &TransfersRepo{db: db}
}

func (TR *TransfersRepo) GetTransfers(ctx context.Context, id int, limit int) ([]db.Transfer, error) {
	query := sq.Select(
		"id",
		"id_from",
		"id_to",
		"currency",
		"value",
		"operation_state",
	).From("transfers").OrderBy("created_at desc").Where(sq.Or{sq.Eq{"id_from": id}, sq.Eq{"id_to": id}})
	if limit != 0 {
		query = query.Limit(uint64(limit))
	}
	query = query.PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := TR.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var items []db.Transfer
	for rows.Next() {
		var item db.Transfer
		if err = rows.Scan(&item.ID, &item.IDFrom, &item.IDTo, &item.Currency, &item.Value, &item.OperationState); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (TR *TransfersRepo) GetTransferByID(ctx context.Context, id int) (*db.Transfer, error) {
	query := sq.Select(
		"id",
		"id_from",
		"id_to",
		"currency",
		"value",
		"operation_state",
	).From("transfers").OrderBy("created_at desc").Where(sq.Eq{"id": id})
	query = query.PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	row := TR.db.QueryRow(ctx, sql, args...)

	item := &db.Transfer{}
	err = row.Scan(&item.ID, &item.IDFrom, &item.IDTo, &item.Currency, &item.Value, &item.OperationState)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, db.NotFound
		}
		return nil, err
	}
	return item, nil
}

func (TR *TransfersRepo) MakeTransfer(ctx context.Context, t *db.Transfer) (int, error) {
	createAt := time.Now()
	sql, args, err := sq.Insert("transfers").
		SetMap(map[string]interface{}{
			"id_from":         t.IDFrom,
			"id_to":           t.IDTo,
			"currency":        t.Currency,
			"value":           t.Value,
			"operation_state": db.Success,
			"created_at":      createAt,
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
		row = TR.db.QueryRow(ctx, sql, args...)
	}

	var id int
	if err = row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (TR *TransfersRepo) CancelTransfer(ctx context.Context, id int) error {
	cancel := sq.Update("transfers").
		Set("operation_state", sq.Expr("?", db.Canceled)).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := cancel.ToSql()
	if err != nil {
		return err
	}

	_, err = TR.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}
