package cli

import (
	"github.com/urfave/cli/v3"

	"github.com/PiskarevSA/minimarket/microservices/points/internal/config"
	pkgcli "github.com/PiskarevSA/minimarket/pkg/cli"
)

func flags() []cli.Flag {
	flags := []cli.Flag{
		pkgcli.ServiceNameFlag,
		pkgcli.PostgreSqlAddrFlag,
		pkgcli.PostgreSqlUserFlag,
		pkgcli.PostgreSqlPasswordFlag,
		pkgcli.PostgreSqlSslModeFlag,
	}

	pkgcli.ServiceNameFlag.Value = "minimarket-points"
	pkgcli.ServiceNameFlag.Destination = &config.Config().ServiceName

	pkgcli.PostgreSqlAddrFlag.Destination = &config.Config().PostgreSqlAddr
	pkgcli.PostgreSqlUserFlag.Destination = &config.Config().PostgreSqlUser
	pkgcli.PostgreSqlPasswordFlag.Destination = &config.Config().PostgreSqlPassword
	pkgcli.PostgreSqlDbFlag.Destination = &config.Config().PostgreSqlDb
	pkgcli.PostgreSqlSslModeFlag.Destination = &config.Config().PostgreSqlSslMode

	return flags
}
