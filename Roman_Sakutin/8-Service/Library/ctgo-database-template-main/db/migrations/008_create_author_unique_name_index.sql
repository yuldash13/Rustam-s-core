-- +goose Up
CREATE UNIQUE INDEX author_name_unique_idx ON author (name);

-- +goose Down
DROP INDEX IF EXISTS author_name_unique_idx;
