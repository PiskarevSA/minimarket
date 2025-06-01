-- name: GetBalanceByUserID :one
SELECT
    current,
    withdrawn
FROM balances
WHERE user_id = @user_id;
