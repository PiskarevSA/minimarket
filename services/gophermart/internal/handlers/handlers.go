package handlers

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/github.com/PiskarevSA/minimarket/pkg/middlewares/jwtauth"
)

var getJwtFromContext = func(ctx context.Context, op string) (token *jwt.Token, ok bool) {
	token, ok = jwtauth.JwtCtxKey.ValueOk(ctx)
	if !ok {
		log.Error().
			Str("op", op).
			Str("layer", "handler").
			Msg("failed to get jwt from ctx")

		return nil, false
	}

	return token, true
}

var getUserIDFromJwt = func(token *jwt.Token, op string) (userId uuid.UUID, ok bool) {
	sub, err := token.Claims.GetSubject()
	if err != nil {
		log.Error().
			Str("op", op).
			Str("layer", "handler").
			Msg("failed to get user id from jwt")

		return uuid.Nil, false
	}

	userId, err = uuid.Parse(sub)
	if err != nil {
		log.Error().
			Err(err).
			Str("op", op).
			Str("layer", "handler").
			Msg("failed to parse sub to user id")

		return uuid.Nil, false

	}

	return userId, true
}
