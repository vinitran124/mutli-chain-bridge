-- +goose Up
INSERT INTO tokens(id,address, chain_id, name, created_at, updated_at) VALUES (gen_random_uuid(),'0x596c14ba2b6e759c73895ce13e64e49054181a7f', '5555', 'VINI', current_timestamp, current_timestamp);

-- +goose Down

