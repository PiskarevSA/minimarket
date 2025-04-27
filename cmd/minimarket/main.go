package main

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"

	authcli "github.com/PiskarevSA/minimarket/microservices/auth/cli"
)

func init() {
	zerolog.LevelFieldName = "lvl"
	zerolog.ErrorFieldName = "err"
	zerolog.MessageFieldName = "msg"
	zerolog.TimeFieldFormat = time.RFC1123

	log.Logger = log.Logger.
		Level(zerolog.InfoLevel).With().
		Timestamp().
		Logger()
}

func main() {
	rootCli := cli.Command{
		Name:    "minimarket",
		Version: "1.0.0",
		Commands: []*cli.Command{
			authcli.AuthCli(),
		},
	}

	err := rootCli.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal().Err(err).
			Msg("failed to setup cli")
	}
}
