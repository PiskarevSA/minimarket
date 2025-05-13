-- name: CreateOutbox :exec
INSERT INTO outbox (
    event,
    payload,
    created_by,
    updated_by
)
SELECT
    @event::VARCHAR(64),
    @payload::JSONB,
    @created_by::VARCHAR(32),
    @updated_by::VARCHAR(32);