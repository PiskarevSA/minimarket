-- name: CreateOrder :one
WITH insertion AS (
    INSERT INTO orders (
        number,
        user_id,
        status,
        accrual,
        uploaded_at
    ) VALUES (
        @number,
        @user_id,
        @status,
        @accrual,
        @uploaded_at
    )
    ON CONFLICT (number) DO NOTHING
    RETURNING user_id, TRUE AS inserted
)
SELECT
  user_id,
  inserted
FROM insertion
UNION ALL
SELECT
  o.user_id,
  FALSE AS inserted
FROM orders o
WHERE o.number = @number
  AND NOT EXISTS (SELECT 1 FROM insertion);