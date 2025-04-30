package cli

import (
	"context"
	"net/http"
	"time"

	"github.com/PiskarevSA/minimarket/microservices/auth/internal/config"
	"github.com/PiskarevSA/minimarket/microservices/auth/internal/oapiserver"
	"github.com/PiskarevSA/minimarket/microservices/auth/internal/oapiserver/oapi"
	"github.com/PiskarevSA/minimarket/microservices/auth/internal/storage/postgre"
	"github.com/PiskarevSA/minimarket/microservices/auth/internal/usecases"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

var Cli = &cli.Command{
	Name:  "auth",
	Usage: "Run 'auth' service",
	Flags: flags,
	Action: func(ctx context.Context, cmd *cli.Command) error {
		userStorage := postgre.NewUser()
		userUsecase := usecases.NewUser(userStorage)

		oapiServer := oapiserver.New(userUsecase)
		router := chi.NewRouter()

		strictHandler := oapi.NewStrictHandler(oapiServer, nil)
		handler := oapi.HandlerFromMux(strictHandler, router)

		server := http.Server{
			Addr:         config.ServerAddr(),
			Handler:      handler,
			ReadTimeout:  3 * time.Second,
			WriteTimeout: 3 * time.Second,
			IdleTimeout:  15 * time.Second,
		}

		log.Info().Str("addr", config.ServerAddr()).
			Msg("listening server...")

		go func() {
			err := server.ListenAndServe()
			if err != nil && err != http.ErrServerClosed {
				log.Fatal().Err(err).
					Msg("failed to listen server")
			}
		}()

		<-ctx.Done()

		return nil
	},
}
