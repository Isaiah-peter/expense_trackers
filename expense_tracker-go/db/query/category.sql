-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1 LIMIT 1;

-- name: ListCategories :many
SELECT * FROM categories
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: ListCategoryByUserId :many
SELECT * FROM categories
WHERE user_id = $1
ORDER BY name
LIMIT $2
OFFSET $3;

-- name: CreateCategory :one
INSERT INTO categories (
  user_id, icon, name
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;

-- name: UpdateCategory :exec
UPDATE categories
SET (name, icon, updated_at) = ($2, $3, $4)
WHERE id = $1;