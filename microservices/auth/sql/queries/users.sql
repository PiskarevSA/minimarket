-- name: CreateUser :exec
INSERT INTO users (id, login, password_hash)
VALUES (@id, @login, @password_hash);

-- name: GetUserCreds :one
SELECT
    id AS user_id,
    password_hash
FROM users
WHERE
    login = @login;
