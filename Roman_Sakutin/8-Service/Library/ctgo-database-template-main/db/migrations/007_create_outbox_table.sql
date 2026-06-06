-- +goose Up
CREATE TYPE outbox_status AS ENUM ('CREATED', 'IN_PROGRESS', 'SUCCESS');

CREATE TABLE outbox
(
    idempotency_key TEXT PRIMARY KEY,
    data            JSONB                   NOT NULL,
    status          outbox_status           NOT NULL,
    kind            INT                     NOT NULL,
    created_at      TIMESTAMP DEFAULT now() NOT NULL,
    updated_at      TIMESTAMP DEFAULT now() NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS outbox;
DROP TYPE IF EXISTS outbox_status;