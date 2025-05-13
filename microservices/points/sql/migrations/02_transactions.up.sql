
CREATE TABLE IF NOT EXISTS transactions (
    id INTEGER NOT NULL,
    order_id UUID,
    user_id UUID NOT NULL,
    operation VARCHAR(16) NOT NULL,
    amount DECIMAL NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    PRIMARY KEY(id, order_id, timestamp)
);

CREATE INDEX IF NOT EXISTS transactions_tx_by_order_idx
    ON transactions (order_id)
    INCLUDE (user_id, timestamp, tx_type, amount);

SELECT create_hypertable(
    'transactions',
    'timestamp',
    chunk_time_interval => INTERVAL '6 days',
    if_not_exists => TRUE
);

CREATE TABLE IF NOT EXISTS transactions_counter (
    user_id UUID PRIMARY KEY,
    counter INTEGER NOT NULL
);
