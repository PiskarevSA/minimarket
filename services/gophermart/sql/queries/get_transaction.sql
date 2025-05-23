-- name: GetTransactionsByUserId :many
SELECT
    id,
    order_number,
    sum,
    processed_at
FROM transactions
WHERE user_id = @user_id AND operation = @operation::VARCHAR(16)
ORDER BY id DESC
OFFSET sqlc.narg('offset')::INTEGER
LIMIT sqlc.narg('limit')::INTEGER;
