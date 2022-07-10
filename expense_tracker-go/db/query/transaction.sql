-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = $1 LIMIT 1;

-- name: ListTransactions :many
SELECT * FROM transactions
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: ListTransactionsByUserId :many
SELECT * FROM transactions
WHERE user_id = $1
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: ListTransactionByStatus :many
SELECT SUM(amount) FROM transactions
WHERE user_id = $1
AND status = $2;

-- name: GetTotalAmount :many
SELECT SUM(ammout) from transactions
WHERE user_id  = $1
AND status = $2
AND status = $3;


-- name: ListTransactionsByCategoryID :many
SELECT * FROM transactions
WHERE category_id = $1
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: CreateTransaction :one
INSERT INTO transactions (
  user_id, category_id, ammout, notes, status
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1;

-- name: UpdateTransaction :exec
UPDATE transactions
SET (ammout, notes, category_id, updated_at) = ($2, $3, $4, $5)
WHERE id = $1;