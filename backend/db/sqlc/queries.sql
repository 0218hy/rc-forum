-- name: ListProducts :many
SELECT * FROM products;

-- name: FindProductByID :one
SELECT * FROM products WHERE id = $1;


-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

