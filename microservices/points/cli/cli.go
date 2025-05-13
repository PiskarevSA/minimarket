package cli

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
	"github.com/urfave/cli/v3"

	"github.com/PiskarevSA/minimarket/microservices/points/internal/config"
	"github.com/PiskarevSA/minimarket/microservices/points/internal/events"
	"github.com/PiskarevSA/minimarket/microservices/points/internal/storage"
	"github.com/PiskarevSA/minimarket/microservices/points/internal/usecases"
)

var Cli = &cli.Command{
	Name:  "points",
	Usage: "Run 'points' service",
	Flags: flags(),
	Action: func(ctx context.Context, c *cli.Command) error {
		pgxPool, err := pgxpool.New(ctx, config.PostgreSqlConnUrl())
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("failed to connect pgxpool")
		}

		storage := storage.New(pgxPool)

		userId := uuid.MustParse("2b5d947e-0f46-47c1-b641-4ac885164b36")
		amount := decimal.NewFromFloat(rand.Float64() * (150 - 10))

		adjustBalance := usecases.NewAdjustBalance(
			config.ServiceName(),
			storage,
		)
		err = adjustBalance.Do(
			ctx,
			userId.String(),
			uuid.NewString(),
			"DEPOSIT",
			amount.String(),
			events.BalanceDeposited,
		)

		fmt.Println(err)

		// balance, err := storage.GetBalance(ctx, userId)
		getBalance := usecases.NewGetBalance(storage)
		balance, _ := getBalance.Do(ctx, userId.String())

		fmt.Println(balance.Available())

		getTransactions := usecases.NewGetTransactions(storage)
		fmt.Println(getTransactions.Do(ctx, userId.String(), 0, 15))

		// storage := storage.New(pgxPool)
		// adjustBalance := usecases.NewAdjustBalance(
		// 	config.ServiceName(),
		// 	storage,
		// )

		// err = adjustBalance.Do(
		// 	ctx,
		// 	orderId, userId,
		// 	objects.TransactionWithdraw, sum,
		// )

		// getBalance := usecases.NewGetBalance(storage)
		// balance, err := getBalance.Do(ctx, userId)

		// fmt.Println(balance, err)

		// getTransactions := usecases.NewGetTransactions(storage)
		// txs, err := getTransactions.Do(ctx, userId, 5, 5)

		// fmt.Println(txs, err)
		return nil
	},
}
