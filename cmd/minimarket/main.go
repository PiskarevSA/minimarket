package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"

	authcli "github.com/PiskarevSA/minimarket/microservices/auth/cli"
	pointscli "github.com/PiskarevSA/minimarket/microservices/points/cli"
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
	rootCtx := context.Background()
	signals := []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL}

	stopCtx, stop := signal.NotifyContext(rootCtx, signals...)
	defer stop()

	rootCli := cli.Command{
		Name:    "minimarket",
		Version: "1.0.0",
		Commands: []*cli.Command{
			authcli.Cli,
			pointscli.Cli,
		},
	}

	err := rootCli.Run(stopCtx, os.Args)
	if err != nil {
		log.Fatal().Err(err).
			Msg("failed to setup cli")
	}
}
