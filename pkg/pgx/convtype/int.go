package convtype

import (
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
)

var ErrIn4IsNull = errors.New("int4 is null")

func Int32ToInt4(i32 int32) pgtype.Int4 {
	return pgtype.Int4{
		Int32: i32,
		Valid: true,
	}
}

func Int4ToInt32(i4 pgtype.Int4) (int32, error) {
	if !i4.Valid {
		return 0, ErrIn4IsNull
	}

	return i4.Int32, nil
}
