-- name: CreateTransaction :exec
WITH next_tx_id AS (
    INSERT INTO transaction_counter (
        user_id,
        counter
    )
    VALUES (
        @user_id::UUID,
        1
    )
    ON CONFLICT (user_id) DO UPDATE
    SET counter = transaction_counter.counter + 1
    RETURNING counter AS id
)
INSERT INTO transactions (
    id,
    user_id,
    order_number,
    operation,
    sum,
    processed_at
)
SELECT
    next_tx_id.id,
    @user_id,
    @order_number,
    @operation,
    @sum,
    @processed_at
FROM next_tx_id;
