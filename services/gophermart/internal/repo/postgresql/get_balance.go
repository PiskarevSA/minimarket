package postgresql

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/dto"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/repo"
)

func (r *PostgreSQL) GetBalanceByUserID(
	ctx context.Context,
	userID objects.UserID,
) (current, withdrawn objects.Amount, err error) {
	userUUID := userID.UUID()

	row, err := r.querier.GetBalanceByUserID(ctx, userUUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = repo.ErrNoBalanceFound
		}

		return objects.NullAmount, objects.NullAmount, err
	}

	current, withdrawn = dto.GetBalanceByUserIDToBalance(row)

	return current, withdrawn, err
}
