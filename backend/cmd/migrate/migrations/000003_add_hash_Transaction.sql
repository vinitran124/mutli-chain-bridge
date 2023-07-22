-- +goose Up
ALTER TABLE transactions ADD COLUMN hash text NOT NULL;
CREATE UNIQUE INDEX ON transactions(hash);


-- +goose Down
ALTER TABLE transactions DROP COLUMN hash;

