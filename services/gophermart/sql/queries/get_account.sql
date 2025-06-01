-- name: GetAccountByUserID :one
SELECT 
    login, 
    password_hash, 
    created_at, 
    updated_at
FROM accounts
WHERE id = @id;

-- name: GetAccountByLogin :one
SELECT 
    id, 
    password_hash, 
    created_at, 
    updated_at
FROM accounts
WHERE login = @login;
