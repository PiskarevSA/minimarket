SELECT remove_retention_policy('transactions');
DROP TABLE transactions;
DROP INDEX transactions_transaction_by_user_idx;
DROP TABLE transactions_counter;
