--- name: CountUsers :one
SELECT COUNT(*)
FROM users
WHERE
    deleted = TRUE;