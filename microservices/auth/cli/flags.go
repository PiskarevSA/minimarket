package cli

import (
	"github.com/urfave/cli/v3"

	"github.com/PiskarevSA/minimarket/microservices/auth/internal/config"
)

var flags = []cli.Flag{
	logLevelFlag,
	serverAddrFlag,
	postgreSqlAddrFlag,
	postgreSqlUserFlag,
	postgreSqlPasswordFlag,
	postgreSqlDbFlag,
	postgreSqlSslModeFlag,
}

var (
	logLevelFlag = &cli.StringFlag{
		Name:        "log.level",
		Value:       "info",
		Destination: &config.Config().LogLevel,
	}

	serverAddrFlag = &cli.StringFlag{
		Name:        "server.addr",
		Value:       "127.0.0.1:8461",
		Destination: &config.Config().ServerAddr,
	}

	postgreSqlAddrFlag = &cli.StringFlag{
		Name:        "postgresql.addr",
		Value:       "127.0.0.1:5432",
		Destination: &config.Config().PostgreSqlAddr,
	}

	postgreSqlUserFlag = &cli.StringFlag{
		Name:        "postgresql.user",
		Value:       "user",
		Destination: &config.Config().PostgreSqlUser,
	}

	postgreSqlPasswordFlag = &cli.StringFlag{
		Name:        "postgresql.password",
		Value:       "password",
		Destination: &config.Config().PostgreSqlPassword,
	}

	postgreSqlDbFlag = &cli.StringFlag{
		Name:        "postgresql.db",
		Value:       "postgres",
		Destination: &config.Config().PostgreSqlDb,
	}

	postgreSqlSslModeFlag = &cli.BoolFlag{
		Name:        "postgresql.sslmode",
		Value:       false,
		Destination: &config.Config().PostgreSqlSslMode,
	}
)
