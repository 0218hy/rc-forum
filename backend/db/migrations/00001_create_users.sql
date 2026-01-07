-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email CITEXT  UNIQUE NOT NULL CHECK (email ~ '^e\d{7}@u\.nus\.edu$'),
    password TEXT NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT false
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
