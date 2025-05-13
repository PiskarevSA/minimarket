package convtype

import (
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

var (
	ErrNumericIsNull     = errors.New("numeric is null")
	ErrNumericIsNan      = errors.New("numeric is nan")
	ErrNumericIsInfinite = errors.New("numeric is infinite")
)

func NumericToDecimal(numeric pgtype.Numeric) (decimal.Decimal, error) {
	if !numeric.Valid {
		return decimal.Decimal{}, ErrNumericIsNull
	}

	if numeric.NaN {
		return decimal.Decimal{}, ErrNumericIsNan
	}

	if numeric.InfinityModifier != pgtype.Finite {
		return decimal.Decimal{}, ErrNumericIsInfinite
	}

	return decimal.NewFromBigInt(numeric.Int, numeric.Exp), nil
}

func DecimalToNumeric(decimal decimal.Decimal) pgtype.Numeric {
	return pgtype.Numeric{
		Int:   decimal.Coefficient(),
		Exp:   decimal.Exponent(),
		Valid: true,
	}
}
