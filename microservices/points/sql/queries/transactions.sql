-- name: CreateTransaction :exec
WITH next_tx_id AS (
    INSERT INTO transactions_counter (
        user_id,
        counter
    )
    VALUES (
        @user_id::UUID,
        1
    )
    ON CONFLICT (user_id) DO UPDATE
    SET counter = transactions_counter.counter + 1
    RETURNING counter AS id
)
INSERT INTO transactions (
    id,
    order_id,
    user_id,
    operation,
    amount
)
SELECT
    next_tx_id.id,
    @order_id::UUID,
    @user_id,
    @operation::VARCHAR(16),
    @amount::DECIMAL
FROM next_tx_id;


-- name: GetTransactions :many
SELECT
    id,
    order_id,
    user_id,
    operation,
    amount,
    timestamp
FROM transactions
WHERE user_id = @user_id::UUID
ORDER BY id DESC
OFFSET sqlc.narg('offset')::INTEGER
LIMIT sqlc.narg('limit')::INTEGER;