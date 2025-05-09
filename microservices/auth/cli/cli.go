package cli

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"

	"github.com/PiskarevSA/minimarket/microservices/auth/internal/config"
	"github.com/PiskarevSA/minimarket/microservices/auth/internal/oapiserver"
	"github.com/PiskarevSA/minimarket/microservices/auth/internal/oapiserver/oapi"
	"github.com/PiskarevSA/minimarket/microservices/auth/internal/storage/postgresql"
	"github.com/PiskarevSA/minimarket/microservices/auth/internal/usecases"
)

var Cli = &cli.Command{
	Name:  "auth",
	Usage: "Run 'auth' service",
	Flags: flags,
	Action: func(ctx context.Context, c *cli.Command) error {
		pool, err := pgxpool.New(ctx, config.PostgreSqlConnUrl())
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		userStorage := postgresql.NewUser(pool)

		jwtSignKeyBytes, err := os.ReadFile(config.Config().JwtSignKeyFilePath)
		if err != nil {
			log.Fatal().Err(err).Send()
		}
		jwtSignKey := string(jwtSignKeyBytes)

		userRegister := usecases.NewUserRegister(
			userStorage,
			jwtSignKey,
			time.Hour,
			24*time.Hour,
		)
		userLogIn := usecases.NewUserLogIn(
			userStorage,
			jwtSignKey,
			time.Hour,
			24*time.Hour,
		)

		router := chi.NewRouter()
		oapiserver := oapiserver.New(userRegister, userLogIn)

		strictHandler := oapi.NewStrictHandler(oapiserver, nil)
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
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatal().Err(err).
					Msg("failed to listen server")
			}
		}()

		<-ctx.Done()

		log.Info().Msg("server stopped")

		return nil
	},
}
