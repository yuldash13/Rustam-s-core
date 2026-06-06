package repository

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/project/library/internal/entity"
)

func mapPostgresErr(err error) error {
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return err
	}

	switch pgErr.Code {
	case "23503":
		return entity.NotFound
	case "23505":
		return entity.Exist
	case "22P02":
		return entity.InvalidArgument
	default:
		return err
	}
}
