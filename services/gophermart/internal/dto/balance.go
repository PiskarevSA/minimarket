package dto

import (
	"github.com/github.com/PiskarevSA/minimarket/pkg/convtype"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/gen/sqlc/postgresql"
)

func GetBalanceByUserIDToBalance(
	row postgresql.GetBalanceByUserIDRow,
) (current, withdrawn objects.Amount) {
	dec, _ := convtype.NumericToDecimal(row.Current)
	current = objects.Amount(dec)

	dec, _ = convtype.NumericToDecimal(row.Withdrawn)
	withdrawn = objects.Amount(dec)

	return current, withdrawn
}
