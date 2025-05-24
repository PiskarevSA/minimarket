package cli

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddlewares "github.com/go-chi/chi/v5/middleware"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"

	pkgmiddlewares "github.com/github.com/PiskarevSA/minimarket/pkg/middlewares"
	"github.com/github.com/PiskarevSA/minimarket/pkg/middlewares/jwtauth"
	"github.com/github.com/PiskarevSA/minimarket/pkg/pgx/transactor"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/config"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/handlers"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/idp"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/repo/postgresql"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/usecases"
)

var Run = &cli.Command{
	Name:  "run",
	Usage: "Запуск gophermart сервиса",
	Flags: runFlags,
	Action: func(ctx context.Context, c *cli.Command) error {
		const serviceName = "auth"

		connURL := config.Config.Database.ConnURL()
		pgxPool, err := pgxpool.New(ctx, connURL)
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("failed to parse postgresql conn url")
		}

		postgreSQL := postgresql.New(pgxPool)
		idp := idp.NewIdentityProvider(
			serviceName,
			jwt.SigningMethodHS256,
			config.Config.Jwt.SigningKey(),
			time.Hour,
		)

		registerUsecase := usecases.NewRegister(postgreSQL, idp)
		loginUsecase := usecases.NewLogin(postgreSQL, idp)

		registerHandler := handlers.NewRegister(registerUsecase)
		loginHandler := handlers.NewLogin(loginUsecase)

		router := chi.NewRouter()

		router.Use(chimiddlewares.Recoverer)
		router.Use(pkgmiddlewares.Decompress)

		registerHandler.Mount(router)
		loginHandler.Mount(router)

		router.Group(func(authRouter chi.Router) {
			ja := jwtauth.New(config.Config.Jwt.SigningKey())
			extractor := jwtauth.ExtractFromAuthHeader
			authRouter.Use(jwtauth.Authenticate(ja, extractor))

			uploadOrderNumberUsecase := usecases.NewUploadOrderNumber(
				postgreSQL,
			)
			getOrdersUsecase := usecases.NewGetOrders(postgreSQL)
			getBalanceUsecase := usecases.NewGetBalance(postgreSQL)
			withdrawUsecase := usecases.NewWithdraw(
				postgreSQL,
				transactor.New(pgxPool),
			)
			getWithdrawalsUsecase := usecases.NewGetWithdrawals(postgreSQL)

			uploadOrderNumberHandler := handlers.NewUploadOrderNumber(
				uploadOrderNumberUsecase,
			)
			getOrdersHandler := handlers.NewGetOrders(getOrdersUsecase)
			getBalanceHandler := handlers.NewGetBalance(getBalanceUsecase)
			withdrawHandler := handlers.NewWithdraw(withdrawUsecase)
			getWithdrawalsHandler := handlers.NewGetWithdrawals(
				getWithdrawalsUsecase,
			)

			getOrdersHandler.Mount(authRouter)
			uploadOrderNumberHandler.Mount(authRouter)
			getBalanceHandler.Mount(authRouter)
			withdrawHandler.Mount(authRouter)
			getWithdrawalsHandler.Mount(authRouter)
		})

		server := http.Server{
			Addr:         config.Config.Server.Addr,
			Handler:      router,
			ReadTimeout:  config.Config.Server.ReadTimeout,
			WriteTimeout: config.Config.Server.WriteTimeout,
			IdleTimeout:  config.Config.Server.IdleTimeout,
		}

		log.Info().
			Str("addr", config.Config.Server.Addr).
			Msg("listening server...")

		go func() {
			err := server.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatal().
					Err(err).
					Msg("failed to run server")
			}
		}()

		<-ctx.Done()

		return nil
	},
}
