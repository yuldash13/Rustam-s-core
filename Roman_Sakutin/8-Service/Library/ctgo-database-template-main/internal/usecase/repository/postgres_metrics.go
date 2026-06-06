package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

func init() {
	prometheus.MustRegister(postgresTableRows)
}

var postgresTableRows = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "postgres_table_rows",
		Help: "Current row count per table",
	},
	[]string{"table"},
)

type TableRowCount struct {
	Table string
	Count int64
}

func (p *postgresRepository) CountTableRows(ctx context.Context) ([]TableRowCount, error) {
	const q = `
SELECT 'author' AS table_name, count(*)::bigint FROM author
UNION ALL
SELECT 'book', count(*)::bigint FROM book
UNION ALL
SELECT 'author_book', count(*)::bigint FROM author_book
UNION ALL
SELECT 'outbox', count(*)::bigint FROM outbox
`
	rows, err := p.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []TableRowCount
	for rows.Next() {
		var m TableRowCount
		if err := rows.Scan(&m.Table, &m.Count); err != nil {
			return nil, err
		}
		out = append(out, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func RunPostgresMetricsCollector(ctx context.Context, pool *pgxpool.Pool, interval time.Duration, logger *zap.Logger) {
	repo := NewPostgresRepository(pool)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	tables := []string{"author", "book", "author_book", "outbox"}

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			collectCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			counts, err := repo.CountTableRows(collectCtx)
			cancel()
			if err != nil {
				logger.Error("postgres metrics collect", zap.Error(err))
				continue
			}

			for _, t := range tables {
				postgresTableRows.WithLabelValues(t).Set(0)
			}
			for _, c := range counts {
				postgresTableRows.WithLabelValues(c.Table).Set(float64(c.Count))
			}
		}
	}
}
