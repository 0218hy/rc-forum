-- name: ListProducts :many
SELECT * FROM products;

-- name: FindProductByID :one
SELECT * FROM products WHERE id = $1;


-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetAllPosts :many
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



