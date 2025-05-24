package objects

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"

	"github.com/github.com/PiskarevSA/minimarket/pkg/convtype"
)

type Amount decimal.Decimal

type AmountError Error

func (e AmountError) Error() string {
	return e.Message
}

var NullAmount = Amount{}

var (
	ErrInvaliAmountValue   = &AmountError{"amount must be non-negative"}
	ErrInvalidAmountFormat = &AmountError{"invalid amount format"}
	ErrInvaliAmountType    = &AmountError{"invalid amount type"}
)

func NewAmount(value any) (Amount, error) {
	var (
		dec decimal.Decimal
		err error
	)

	switch value := any(value).(type) {
	case string:
		dec, err = decimal.NewFromString(value)
		if err != nil {
			return NullAmount, ErrInvalidAmountFormat
		}

	case pgtype.Numeric:
		dec, err = convtype.NumericToDecimal(value)
		if err != nil {
			return NullAmount, err
		}
	default:
		return NullAmount, ErrInvaliAmountType
	}

	if dec.LessThan(decimal.Zero) {
		return NullAmount, ErrInvaliAmountValue
	}

	return Amount(dec), nil
}

func (a Amount) Decimal() decimal.Decimal {
	return decimal.Decimal(a)
}

func (a Amount) Numeric() pgtype.Numeric {
	dec := a.Decimal()

	return convtype.DecimalToNumeric(dec)
}

func (a Amount) String() string {
	return a.Decimal().String()
}

func (a Amount) Equal(other Amount) bool {
	dec := decimal.Decimal(a)

	return dec.Equal(decimal.Decimal(other))
}
