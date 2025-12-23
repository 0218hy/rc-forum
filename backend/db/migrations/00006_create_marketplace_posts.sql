-- +goose Up
-- +goose StatementBegin
CREATE TYPE listing_type AS ENUM ('buy', 'sell', 'give');
CREATE TYPE listing_status_type AS ENUM ('open', 'reserved', 'completed');

CREATE TABLE IF NOT EXISTS marketplace_posts (
    post_id INT PRIMARY KEY REFERENCES posts(id) ON DELETE CASCADE,
    listing listing_type NOT NULL,
    price DECIMAL(19,4),
    quantity INT NOT NULL,
    listing_status listing_status_type NOT NULL DEFAULT 'open',
    CONSTRAINT chk_price_not_null_condition CHECK (
        (listing <> 'give') OR (price IS NOT NULL)
    )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS marketplace_posts;
-- +goose StatementEnd
