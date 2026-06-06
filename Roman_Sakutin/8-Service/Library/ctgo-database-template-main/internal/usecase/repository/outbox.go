package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/project/library/internal/entity"
)

const (
	OutboxKindAuthor = 1
	OutboxKindBook   = 2
)

type OutboxRepository interface {
	SendMessage(ctx context.Context, idempotencyKey string, kind int, data []byte) error
}

type outboxRepository struct {
	pool *pgxpool.Pool
}

func NewOutboxRepository(pool *pgxpool.Pool) OutboxRepository {
	return &outboxRepository{pool: pool}
}

func (o *outboxRepository) q(ctx context.Context) PoolTx {
	if tx, ok := TxFromContext(ctx); ok {
		return tx
	}
	return o.pool
}

func (o *outboxRepository) SendMessage(ctx context.Context, idempotencyKey string, kind int, data []byte) error {
	const q = `
INSERT INTO outbox (idempotency_key, data, status, kind)
VALUES ($1, $2, 'CREATED'::outbox_status, $3)
ON CONFLICT (idempotency_key) DO NOTHING
`
	_, err := o.q(ctx).Exec(ctx, q, idempotencyKey, data, kind)
	return err
}

type OutboxMessage struct {
	Key  string
	Data []byte
	Kind int
}

type OutboxStore struct {
	pool *pgxpool.Pool
}

func NewOutboxStore(pool *pgxpool.Pool) *OutboxStore {
	return &OutboxStore{pool: pool}
}

func (s *OutboxStore) ReclaimStale(ctx context.Context, ttl time.Duration) error {
	if ttl <= 0 {
		return nil
	}
	const q = `
UPDATE outbox
SET status = 'CREATED'::outbox_status,
    updated_at = now()
WHERE status = 'IN_PROGRESS'::outbox_status
  AND updated_at < now() - ($1::bigint * interval '1 millisecond')
`
	ms := ttl.Milliseconds()
	if ms < 1 {
		ms = 1
	}
	_, err := s.pool.Exec(ctx, q, ms)
	return err
}

func (s *OutboxStore) ClaimBatch(ctx context.Context, limit int) ([]OutboxMessage, error) {
	if limit < 1 {
		limit = 1
	}
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	const q = `
WITH picked AS (
    SELECT idempotency_key
    FROM outbox
    WHERE status = 'CREATED'::outbox_status
    ORDER BY created_at ASC
    LIMIT $1
    FOR UPDATE SKIP LOCKED
)
UPDATE outbox o
SET status = 'IN_PROGRESS'::outbox_status,
    updated_at = now()
FROM picked p
WHERE o.idempotency_key = p.idempotency_key
RETURNING o.idempotency_key, o.data, o.kind
`
	rows, err := tx.Query(ctx, q, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []OutboxMessage
	for rows.Next() {
		var m OutboxMessage
		if err := rows.Scan(&m.Key, &m.Data, &m.Kind); err != nil {
			return nil, err
		}
		out = append(out, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *OutboxStore) TouchInProgress(ctx context.Context, key string) error {
	const q = `
UPDATE outbox
SET updated_at = now()
WHERE idempotency_key = $1
  AND status = 'IN_PROGRESS'::outbox_status
`
	_, err := s.pool.Exec(ctx, q, key)
	return err
}

func (s *OutboxStore) MarkSuccess(ctx context.Context, key string) error {
	const q = `
UPDATE outbox
SET status = 'SUCCESS'::outbox_status,
    updated_at = now()
WHERE idempotency_key = $1
  AND status = 'IN_PROGRESS'::outbox_status
`
	tag, err := s.pool.Exec(ctx, q, key)
	if err != nil {
		return err
	}
	if tag.RowsAffected() > 0 {
		return nil
	}

	var status string
	err = s.pool.QueryRow(ctx, `SELECT status::text FROM outbox WHERE idempotency_key = $1`, key).Scan(&status)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}
	if status == "SUCCESS" {
		return nil
	}

	return errors.New("outbox: mark success: no row")
}

func (s *OutboxStore) MarkRetry(ctx context.Context, key string) error {
	const q = `
UPDATE outbox
SET status = 'CREATED'::outbox_status,
    updated_at = now()
WHERE idempotency_key = $1
  AND status = 'IN_PROGRESS'::outbox_status
`
	_, err := s.pool.Exec(ctx, q, key)
	return err
}

func (s *OutboxStore) MarkDelivered(ctx context.Context, key string) error {
	const q = `
UPDATE outbox
SET status = 'SUCCESS'::outbox_status,
    updated_at = now()
WHERE idempotency_key = $1
  AND status != 'SUCCESS'::outbox_status
`
	_, err := s.pool.Exec(ctx, q, key)
	return err
}

func (s *OutboxStore) Status(ctx context.Context, key string) (string, error) {
	const q = `
SELECT status::text
FROM outbox
WHERE idempotency_key = $1
`
	var status string
	err := s.pool.QueryRow(ctx, q, key).Scan(
		&status,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "nil", entity.NotFound
		}
		return "nil", mapPostgresErr(err)
	}
	return status, err
}

func (s *OutboxStore) ReclaimAllInProgress(ctx context.Context) error {
	const q = `
UPDATE outbox
SET status = 'CREATED'::outbox_status,
    updated_at = now()
WHERE status = 'IN_PROGRESS'::outbox_status
`
	_, err := s.pool.Exec(ctx, q)
	return err
}

type OutboxStatusCount struct {
	Kind   int
	Status string
	Count  int64
}

func (s *OutboxStore) CountByStatusAndKind(ctx context.Context) ([]OutboxStatusCount, error) {
	const q = `
SELECT kind, status::text, count(*)::bigint
FROM outbox
GROUP BY kind, status
`
	rows, err := s.pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []OutboxStatusCount
	for rows.Next() {
		var m OutboxStatusCount
		if err := rows.Scan(&m.Kind, &m.Status, &m.Count); err != nil {
			return nil, err
		}
		out = append(out, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}
