-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS announcement_posts (
    post_id INT PRIMARY KEY REFERENCES posts(id) ON DELETE CASCADE,
    expires_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS announcement_posts;
-- +goose StatementEnd
