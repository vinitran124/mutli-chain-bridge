-- +goose Up
INSERT INTO tokens( id, address, chain_id, name, created_at, updated_at) VALUES (gen_random_uuid(),'0x596c14ba2b6e759c73895ce13e64e49054181a7f', '5555', 'VINI', current_timestamp, current_timestamp);
INSERT INTO tokens( id, address, chain_id, name, created_at, updated_at) VALUES (gen_random_uuid(),'0x32D7eEE6513224f459D221BfED0C3af45343CbB7', '97', 'VINI', current_timestamp, current_timestamp);

-- +goose Down
DELETE FROM tokens;

