CREATE TABLE IF NOT EXISTS transactions (
    id INTEGER NOT NULL,
    user_id UUID NOT NULL,
    order_number VARCHAR(16) NOT NULL,
    operation VARCHAR(16) NOT NULL,
    sum DECIMAL NOT NULL,
    processed_at TIMESTAMPTZ NOT NULL,

    PRIMARY KEY(id, order_number, processed_at)
);

CREATE TABLE IF NOT EXISTS transaction_counter (
    user_id UUID PRIMARY KEY,
    counter INTEGER NOT NULL
);
