-- +goose Up
ALTER TABLE transactions ADD COLUMN hash text NOT NULL;
-- +goose Down
ALTER TABLE transactions DROP COLUMN hash;

