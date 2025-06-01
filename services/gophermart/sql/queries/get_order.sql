-- name: GetOrderByNumber :one
SELECT 
    user_id, 
    status, 
    accrual, 
    uploaded_at
FROM orders
WHERE number = @number;

-- name: GetOrdersByUserID :many
SELECT 
    number, 
    status, 
    accrual, 
    uploaded_at
FROM orders
WHERE user_id = @user_id
ORDER BY uploaded_at DESC
OFFSET sqlc.narg('offset')::INTEGER
LIMIT sqlc.narg('limit')::INTEGER;
