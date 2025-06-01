-- name: CreateRewardStrategy :exec
INSERT INTO reward_strategies (
    match,
    reward,
    reward_type
) VALUES (
    @match,
    @reward,
    @reward_type
);
