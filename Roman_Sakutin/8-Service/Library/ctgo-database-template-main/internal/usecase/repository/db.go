package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"

	"github.com/project/library/config"
)

func CreateDatabasePoolConnections(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	poolCfg, err := pgxpool.ParseConfig(cfg.PG.DB)
	if err != nil {
		return nil, errors.Wrap(err, "invalid config")
	}

	poolCfg.MaxConns = 10
	poolCfg.MinConns = 2
	poolCfg.MaxConnLifetime = time.Hour
	poolCfg.MaxConnIdleTime = time.Minute * 30
	poolCfg.HealthCheckPeriod = time.Minute

	pool, err := pgxpool.ConnectConfig(ctx, poolCfg)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create connection pool")
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, errors.Wrap(err, "unable to ping database")
	}
	return pool, nil
}
