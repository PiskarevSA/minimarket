CREATE TYPE TRANSACTION_TYPE AS ENUM(
    'DEPOSIT', 
    'WITHDRAW', 
    'CURRENT_BALANCE', 
    'TOTAL_WITHDRAWN'
);

CREATE TABLE IF NOT EXISTS transactions (
    id NUMBER NOT NULL,
    order_id UUID NOT NULL,
    user_id UUID NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    type TRANSACTION_TYPE NOT NULL,

    PRIMARY KEY (id, order_id, timestamp),

    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (user_id) REFERENCES users(id);
)