CREATE TABLE IF NOT EXISTS users (
    id            UUID        PRIMARY KEY,
    login         TEXT        UNIQUE NOT NULL,
    password_hash CHAR(60)    NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
