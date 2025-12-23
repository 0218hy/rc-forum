-- +goose Up
-- +goose StatementBegin
CREATE TYPE activity_category_type AS ENUM (
    'food&drinks',
    'sports&fitness',
    'entertainment',
    'outdoor&nature',
    'study&work',
    'events'
);

CREATE TABLE IF NOT EXISTS openjio_posts (
    post_id INT PRIMARY KEY REFERENCES posts(id) ON DELETE CASCADE,
    activity_category activity_category_type NOT NULL,
    location TEXT NOT NULL,
    event_date DATE NOT NULL,
    start_time TIME,
    end_time TIME
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS openjio_posts;
-- +goose StatementEnd
