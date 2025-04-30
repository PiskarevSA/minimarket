package cli

import (
	"github.com/PiskarevSA/minimarket/microservices/auth/internal/config"
	"github.com/PiskarevSA/minimarket/pkg/valiadtors"
	"github.com/urfave/cli/v3"
)

var (
	logLevelFlag = &cli.StringFlag{
		Name:        "log.level",
		Value:       "info",
		Destination: &config.Config().LogLevel,
		Validator: func(l string) error {
			return valiadtors.ValidateLogLevel(l)
		},
	}

	serverAddrFlag = &cli.StringFlag{
		Name:        "server.addr",
		Value:       "127.0.0.1:8624",
		Destination: &config.Config().ServerAddr,
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
