CREATE TABLE IF NOT EXISTS reward_strategies (
    match VARCHAR(128) PRIMARY KEY,
    reward DECIMAL NOT NULL,
    reward_type VARCHAR(12) NOT NULL
);