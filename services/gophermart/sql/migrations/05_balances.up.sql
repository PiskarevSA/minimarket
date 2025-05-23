CREATE TABLE IF NOT EXISTS balances (
    user_id UUID PRIMARY KEY,
    current DECIMAL NOT NULL CHECK(current >= 0),
    withdrawn DECIMAL NOT NULL
);
