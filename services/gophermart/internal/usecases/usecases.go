package usecases

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/entities"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
)

type (
	IdentityProvider interface {
		IssueToken(
			userId uuid.UUID,
			now time.Time,
		) (tokenString string, err error)
	}

	Transactor interface {
		Transact(
			ctx context.Context,
			opts pgx.TxOptions,
			fn func(ctx context.Context) error,
		) error
	}

	TransactionStorage interface {
		CreateTransactionInTx(
			ctx context.Context,
			tx entities.Transaction,
		) error
	}

	BalanceStorage interface {
		CreateOrUpdateBalanceInTx(
			ctx context.Context,
			userId objects.UserID,
			operation objects.Operation,
			sum objects.Amount,
		) error
	}
)
