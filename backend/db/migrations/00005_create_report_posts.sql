-- +goose Up
-- +goose StatementBegin
CREATE TYPE report_status AS ENUM ('open', 'in_progress', 'resolved');
CREATE TYPE urgency_level AS ENUM ('low', 'medium', 'high');

CREATE TABLE IF NOT EXISTS report_posts (
    post_id INT PRIMARY KEY REFERENCES posts(id) ON DELETE CASCADE,
    status report_status NOT NULL DEFAULT 'open',
    urgency urgency_level NOT NULL DEFAULT 'low'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS report_posts;
-- +goose StatementEnd
