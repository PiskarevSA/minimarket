package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/PiskarevSA/minimarket/microservices/points/internal/domain/objects"
)

type Balance struct {
	userId    uuid.UUID
	available objects.Amount
	withdrawn objects.Amount
	updatedAt time.Time
}

func NewBalance[AmountT string | pgtype.Numeric](
	userId uuid.UUID,
	available AmountT,
	withdrawn AmountT,
	updatedAt time.Time,
) (Balance, error) {
	var (
		balance Balance
		err     error
	)

	balance.available, err = objects.NewAmount(available)
	if err != nil {
		return Balance{}, err
	}

	balance.withdrawn, err = objects.NewAmount(withdrawn)
	if err != nil {
		return Balance{}, err
	}

	balance.userId = userId
	balance.updatedAt = updatedAt

	return balance, nil
}

func (b Balance) UserId() uuid.UUID {
	return b.userId
}

func (b Balance) Available() objects.Amount {
	return b.available
}

func (b Balance) Withdrawn() objects.Amount {
	return b.withdrawn
}

func (b Balance) UpdatedAt() time.Time {
	return b.updatedAt
}
