-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users
SET (name, updated_at) = ($2, $3)
WHERE id = $1;