CREATE TABLE IF NOT EXISTS orders (
    number VARCHAR(12) PRIMARY KEY,
    user_id UUID NOT NULL,
    status VARCHAR(16) NOT NULL,
    accrual DECIMAL,
    uploaded_at TIMESTAMPTZ NOT NULL
);