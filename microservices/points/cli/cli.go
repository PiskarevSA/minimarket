package cli

import (
	"context"

	"github.com/urfave/cli/v3"
)

var pointsCli = &cli.Command{
	Name:  "points",
	Usage: "Run 'points' service",
	Flags: flags,
	Action: func(ctx context.Context, cmd *cli.Command) error {
		return nil
	},
}

func Cli() *cli.Command {
	return pointsCli
}
