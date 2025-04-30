-- name: CreateUser :one
WITH inserted AS (
    INSERT INTO 
        users (
            id, 
            login, 
            password_hash, 
            password_salt
        )
    VALUES 
        (
            @id :: UUID, 
            @login :: TEXT, 
            @password_hash :: TEXT, 
            @password_salt :: TEXT
        )
    ON CONFLICT (login) 
        DO NOTHING
    RETURNING 0 AS status
)

SELECT 
    status
FROM 
    inserted
UNION ALL
SELECT 
    1 AS code
FROM 
    users
WHERE 
    login = @login :: TEXT
    AND NOT EXISTS (SELECT
        1 FROM inserted
    )
LIMIT 1;