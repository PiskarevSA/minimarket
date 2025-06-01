-- name: CreateOrUpdateBalance :exec
WITH balance_exists AS (
    SELECT EXISTS (
        SELECT 1
        FROM balances
        WHERE user_id = @user_id
    ) AS exists_flag
)
INSERT INTO balances (
    user_id,
    current,
    withdrawn
)
SELECT
    @user_id AS user_id,
    CASE
    WHEN @operation::VARCHAR(16) = 'DEPOSIT' THEN @sum
    ELSE (
        CASE
            WHEN (SELECT exists_flag FROM balance_exists) THEN 0
            ELSE -1
        END
    )
    END AS current,
    CASE
        WHEN @operation = 'WITHDRAW' THEN @sum
        ELSE 0
    END AS withdrawn
ON CONFLICT (user_id) DO UPDATE
SET
    current = balances.current + (
    CASE
        WHEN @operation = 'DEPOSIT' THEN @sum
        ELSE -(@sum)
    END
    ),
    withdrawn = balances.withdrawn + (
    CASE
        WHEN @operation = 'WITHDRAW' THEN @sum
        ELSE 0
    END
    );
