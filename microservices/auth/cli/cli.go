package cli

import (
	"context"

	"github.com/urfave/cli/v3"
)

var authCli = &cli.Command{
	Name: "auth",
	Flags: []cli.Flag{
		logLevelFlag, serverAddrFlag,
		postgreAddrFlag, postgreUsernameFlag,
		postgrePasswordFlag,
	},
	Action: func(ctx context.Context, cmd *cli.Command) error {
		return nil
	},
}

func AuthCli() *cli.Command {
	return authCli
}
