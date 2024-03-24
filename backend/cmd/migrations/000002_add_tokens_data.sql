-- +goose Up
INSERT INTO tokens( id, address, chain_id, name, created_at, updated_at) VALUES (gen_random_uuid(),'0x15f8253779428d9ea5b054deef3e454d539ddf7e', '11155111', 'VINI', current_timestamp, current_timestamp);
INSERT INTO tokens( id, address, chain_id, name, created_at, updated_at) VALUES (gen_random_uuid(),'0x6b08b796b4b43d565c34cf4b57d8c871db410ebe', '97', 'VINI', current_timestamp, current_timestamp);

-- +goose Down
DELETE FROM tokens;

