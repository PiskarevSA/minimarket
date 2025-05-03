package validators

import "errors"

const (
	traceLvl = "trace"
	debugLvl = "debug"
	infoLvl  = "info"
	warnLvl  = "warn"
	errorLvl = "error"
	fatalLvl = "fatal"
	panicLvl = "panic"
)

var ErrInvalidLogLevel = errors.New("invalid log level")

func ValidateLogLevel(l string) error {
	if l == traceLvl || l == debugLvl || l == infoLvl ||
		l == warnLvl || l == errorLvl || l == fatalLvl ||
		l == panicLvl {
		return nil
	}

	return ErrInvalidLogLevel
}
