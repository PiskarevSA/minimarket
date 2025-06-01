package handlers

import "context"

type RegisterStrategyUsecase interface {
	Do(
		ctx context.Context,
		reward,
		rewardType string,
	) error
}
