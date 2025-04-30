package cli

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"

	chimiddleware "github.com/go-chi/chi/v5/middleware"
	oapimiddleware "github.com/oapi-codegen/nethttp-middleware"

	"github.com/PiskarevSA/minimarket/microservices/points/internal/config"
	"github.com/PiskarevSA/minimarket/microservices/points/internal/oapiserver"
	"github.com/PiskarevSA/minimarket/microservices/points/internal/oapiserver/oapi"
)

var Cli = &cli.Command{
	Name:  "points",
	Usage: "Run 'points' service",
	Flags: flags,
	Action: func(ctx context.Context, cmd *cli.Command) error {
		oapiserver := oapiserver.New()

		router := chi.NewRouter()
		router.Use(chimiddleware.Recoverer)

		swagger, _ := oapi.GetSwagger()
		swagger.Servers = nil

		validator := oapimiddleware.OapiRequestValidatorWithOptions(swagger,
			&oapimiddleware.Options{Options: openapi3filter.Options{AuthenticationFunc: func(ctx context.Context, ai *openapi3filter.AuthenticationInput) error {
				fmt.Println(ctx.Value(oapi.BearerAuthScopes))
				return nil
			}}})
		router.Use(validator)

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
			if err != nil && err != http.ErrServerClosed {
				log.Fatal().Err(err).
					Msg("failed to listen server")
			}
		}()

		<-ctx.Done()

		return nil
	},
}
