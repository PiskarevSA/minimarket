package damm

import (
	"errors"
	"strings"
)

var dammTable = [10][10]int{
	{0, 3, 1, 7, 5, 9, 8, 6, 4, 2},
	{7, 0, 9, 2, 1, 5, 4, 8, 6, 3},
	{4, 2, 0, 6, 8, 7, 1, 3, 5, 9},
	{1, 7, 5, 0, 9, 8, 3, 4, 2, 6},
	{6, 1, 2, 3, 0, 4, 5, 9, 7, 8},
	{3, 6, 7, 4, 2, 0, 9, 5, 8, 1},
	{5, 8, 6, 9, 7, 2, 0, 1, 3, 4},
	{8, 9, 4, 5, 3, 6, 2, 0, 1, 7},
	{9, 4, 3, 8, 6, 1, 7, 2, 0, 5},
	{2, 5, 8, 1, 4, 3, 6, 7, 9, 0},
}

var (
	ErrEmptyValue   = errors.New("damm: empty value")
	ErrNonDigitChar = errors.New("damm: non-digit char")
)

func encode(value string, vlen int) (int, error) {
	if value == "" {
		return 0, ErrEmptyValue
	}

	check := 0

	for i := range vlen {
		d := value[i] - '0'
		if d > 9 {
			return 0, ErrNonDigitChar
		}

		check = dammTable[check][d]
	}

	return check, nil
}

func Append(value string) (string, error) {
	check, err := encode(value, len(value))
	if err != nil {
		return "", err
	}

	var builder strings.Builder

	builder.WriteString(value)
	builder.WriteByte(byte('0' + check))

	return builder.String(), nil
}

func Verify(value string) (bool, error) {
	vlen := len(value)
	if vlen < 1 {
		return false, ErrEmptyValue
	}

	body := value[:vlen-1]
	checkChar := value[vlen-1]

	check, err := encode(body, len(body))
	if err != nil {
		return false, err
	}

	return byte('0'+check) == checkChar, nil
}
