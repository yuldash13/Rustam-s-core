package repo

import (
	"context"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/domain"
)

type Transfers interface {
	GetTransfers(ctx context.Context, id int64, limit int64) ([]domain.Transfer, error)
	GetTransferByID(ctx context.Context, id int64) (*domain.Transfer, error)
	MakeTransfer(ctx context.Context, t *domain.Transfer) (int64, error)
	CancelTransfer(ctx context.Context, id int64) error
}

type TransfersRepo struct {
	domain *pgxpool.Pool
}

func NewTransferRepo(domain *pgxpool.Pool) *TransfersRepo {
	return &TransfersRepo{domain: domain}
}

func (TR *TransfersRepo) GetTransfers(ctx context.Context, id int64, limit int64) ([]domain.Transfer, error) {
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

	rows, err := TR.domain.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var items []domain.Transfer
	for rows.Next() {
		var item domain.Transfer
		if err = rows.Scan(&item.ID, &item.IDFrom, &item.IDTo, &item.Currency, &item.Value, &item.OperationState); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (TR *TransfersRepo) GetTransferByID(ctx context.Context, id int64) (*domain.Transfer, error) {
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

	row := TR.domain.QueryRow(ctx, sql, args...)

	item := &domain.Transfer{}
	err = row.Scan(&item.ID, &item.IDFrom, &item.IDTo, &item.Currency, &item.Value, &item.OperationState)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.NotFound
		}
		return nil, err
	}
	return item, nil
}

func (TR *TransfersRepo) MakeTransfer(ctx context.Context, t *domain.Transfer) (int64, error) {
	createAt := time.Now()
	sql, args, err := sq.Insert("transfers").
		SetMap(map[string]interface{}{
			"id_from":         t.IDFrom,
			"id_to":           t.IDTo,
			"currency":        t.Currency,
			"value":           t.Value,
			"operation_state": domain.Success,
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
		row = TR.domain.QueryRow(ctx, sql, args...)
	}

	var id int64
	if err = row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (TR *TransfersRepo) CancelTransfer(ctx context.Context, id int64) error {
	cancel := sq.Update("transfers").
		Set("operation_state", sq.Expr("?", domain.Canceled)).
		Where(sq.Eq{"id": id}).Where("operation_state", domain.Success).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := cancel.ToSql()
	if err != nil {
		return domain.NotFound
	}

	_, err = TR.domain.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}
