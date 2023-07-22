-- +goose Up
CREATE TABLE "bridge_requests"(
    "id" uuid PRIMARY KEY,
    "user_address" text NOT NULL UNIQUE,
    "input_chain" text NOT NULL,
    "output_chain" text NOT NULL,
    "raw_amount" text NOT NULL,
    "is_complete" boolean NOT NULL DEFAULT FALSE,
    "token" text NOT NULL,
    "created_at" timestamp with time zone NOT NULL,
    "updated_at" timestamp with time zone NOT NULL
);

CREATE UNIQUE INDEX ON bridge_requests(user_address);

-- +goose Down
DROP TABLE IF EXISTS bridge_requests;
