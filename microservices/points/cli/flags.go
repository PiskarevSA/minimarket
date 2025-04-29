package cli

import (
	"github.com/PiskarevSA/minimarket/pkg/valiadtors"
	"github.com/urfave/cli/v3"
)

var (
	logLevelFlag = &cli.StringFlag{
		Name: "log.level",
		Validator: func(l string) error {
			return valiadtors.ValidateLogLevel(l)
		},
	}

	serverAddrFlag = &cli.StringFlag{
		Name: "server.addr",
		Validator: func(a string) error {
			return valiadtors.ValidateAddr(a)
		},
	}

	postgreAddrFlag = &cli.StringFlag{
		Name: "postgre.addr",
	}

	postgreUsernameFlag = &cli.StringFlag{
		Name: "postgre.username",
	}

	postgrePasswordFlag = &cli.StringFlag{
		Name: "postgre.password",
	}

	jwtSecretFlag = &cli.StringFlag{
		Name: "jwt.secret",
	}
)

var flags = []cli.Flag{
	logLevelFlag,
	serverAddrFlag,
	postgreAddrFlag,
	postgreUsernameFlag,
	postgrePasswordFlag,
	jwtSecretFlag,
}
