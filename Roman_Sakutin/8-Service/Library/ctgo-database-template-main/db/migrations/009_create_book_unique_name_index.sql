-- +goose Up
CREATE UNIQUE INDEX book_name_unique_idx ON book (name);

-- +goose Down
DROP INDEX IF EXISTS book_name_unique_idx;
