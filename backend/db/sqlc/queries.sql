-- name: CreateUser :one
INSERT INTO users (
    name,
    email,
    password,
    is_admin
) VALUES ($1, $2, $3, $4) RETURNING id;

-- name: FindUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: ListAllPosts :many
SELECT
    p.*,
    a.expires_at AS announcement_expires_at,
    r.status        AS report_status,
    r.urgency       AS report_urgency,
    m.listing AS marketplace_listing,
    m.price AS marketplace_price,
    m.quantity AS marketplace_quantity,
    m.listing_status AS marketplace_listing_status,
    o.activity_category AS openjio_activity_category,
    o.location AS openjio_location,
    o.event_date AS openjio_event_date,
    o.start_time AS openjio_start_time,
    o.end_time AS openjio_end_time
FROM posts p
LEFT JOIN announcement_posts a
    ON p.type = 'announcement'
   AND a.post_id = p.id
LEFT JOIN report_posts r
    ON p.type = 'report'
   AND r.post_id = p.id
LEFT JOIN marketplace_posts m 
    ON p.type = 'marketplace'
    AND m.post_id = p.id
LEFT JOIN openjio_posts o
    ON p.type = 'openjio'
   AND o.post_id = p.id
ORDER BY p.created_at DESC;

-- name: FindPostByID :one
SELECT
    p.*,
    a.expires_at AS announcement_expires_at,
    r.status        AS report_status,
    r.urgency       AS report_urgency,
    m.listing AS marketplace_listing,
    m.price AS marketplace_price,
    m.quantity AS marketplace_quantity,
    m.listing_status AS marketplace_listing_status,
    o.activity_category AS openjio_activity_category,
    o.location AS openjio_location,
    o.event_date AS openjio_event_date,
    o.start_time AS openjio_start_time,
    o.end_time AS openjio_end_time
FROM posts p
LEFT JOIN announcement_posts a
    ON p.type = 'announcement'
   AND a.post_id = p.id
LEFT JOIN report_posts r
    ON p.type = 'report'
   AND r.post_id = p.id
LEFT JOIN marketplace_posts m 
    ON p.type = 'marketplace'
    AND m.post_id = p.id
LEFT JOIN openjio_posts o
    ON p.type = 'openjio'
   AND o.post_id = p.id
WHERE p.id = $1;

-- name: DeletePostByID :exec
DELETE FROM posts WHERE id = $1;

-- name: CreatePost :one
INSERT INTO posts (
    author_id,
    type,
    title,
    body
)  VALUES ($1, $2, $3, $4) RETURNING *;

-- name: CreateAnnouncement :exec
INSERT INTO announcement_posts (
    post_id,
    expires_at
) VALUES ($1, $2);

-- name: CreateReport :exec
INSERT INTO report_posts (
    post_id,
    status,
    urgency
) VALUES ($1, $2, $3);

-- name: CreateMarketplace :exec
INSERT INTO marketplace_posts (
    post_id,
    listing,
    price,
    quantity,
    listing_status
) VALUES ($1, $2, $3, $4, $5);

-- name: CreateOpenjio :exec
INSERT INTO openjio_posts (
    post_id,
    activity_category,
    location,
    event_date,
    start_time,
    end_time
) VALUES ($1, $2, $3, $4, $5, $6);

-- name: UpdatePostCore :exec
UPDATE posts
SET title = COALESCE($2, title),
    body = COALESCE($3, body),
    updated_at = NOW()
WHERE id = $1;


-- name: GetSession :one
SELECT * FROM sessions WHERE id = $1;

-- name: CreateSession :one
INSERT INTO sessions (
    user_id,
    refresh_token,
    is_revoked,
    expires_at
) VALUES ($1, $2, $3, $4) RETURNING *;  

-- name: RevokeSession :exec
UPDATE sessions SET is_revoked = true WHERE id = $1;    

-- name: DeleteSessionsByUserID :exec
DELETE FROM sessions WHERE user_id = $1;  

-- name: DeleteUserByID :exec
DELETE FROM users WHERE id = $1;

-- name: CreateComment :one
INSERT INTO comments (
    post_id,
    author_id,
    body
) VALUES ($1, $2, $3) RETURNING *;

-- name: ListCommentsByPostID :many
SELECT * FROM comments WHERE post_id = $1 ORDER BY created_at ASC;  

-- name: DeleteCommentByID :exec
DELETE FROM comments WHERE id = $1 AND author_id = $2;

-- name: GetPostsWithAuthors :many
SELECT
  p.id,
  p.type,
  p.body,
  p.created_at,
  u.id   AS author_id,
  u.name AS author_name
FROM posts p
JOIN users u ON p.author_id = u.id
ORDER BY p.created_at DESC;

-- name: GetCommentsByPostIDs :many
SELECT
  c.id,
  c.post_id,
  u.name AS author,
  c.body,
  c.created_at
FROM comments c
JOIN users u ON c.author_id = u.id
WHERE c.post_id = ANY($1::int[])
ORDER BY c.created_at ASC;
