package cli

import (
	"github.com/urfave/cli/v3"
)

var (
	ServiceNameFlag = &cli.StringFlag{
		Name: "service.name",
	}

	PostgreSqlAddrFlag = &cli.StringFlag{
		Name:  "postgresql.addr",
		Value: "127.0.0.1:5432",
	}

	PostgreSqlUserFlag = &cli.StringFlag{
		Name:  "postgresql.user",
		Value: "user",
	}

	PostgreSqlPasswordFlag = &cli.StringFlag{
		Name:  "postgresql.password",
		Value: "password",
	}

	PostgreSqlDbFlag = &cli.StringFlag{
		Name:  "postgresql.db",
		Value: "postgres",
	}

	PostgreSqlSslModeFlag = &cli.BoolFlag{
		Name:  "postgresql.sslmode",
		Value: false,
	}
)
