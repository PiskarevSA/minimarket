package cli

import (
	"context"

	"github.com/urfave/cli/v3"
)

var Cli = &cli.Command{
	Name:  "auth",
	Usage: "Run 'auth' service",
	Flags: flags,
	Action: func(ctx context.Context, cmd *cli.Command) error {
		return nil
	},
}
