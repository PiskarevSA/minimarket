package cli

import "github.com/urfave/cli/v3"

var (
	logLevelFlag = &cli.StringFlag{
		Name: "log.level",
	}

	serverAddrFlag = &cli.StringFlag{
		Name: "server.addr",
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
)
