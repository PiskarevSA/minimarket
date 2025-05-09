package jwtauth

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type (
	ParseFn     func(context.Context, *JwtAuth) jwt.Keyfunc
	ValidatorFn func(context.Context, *jwt.Token) error
)

type JwtAuth struct {
	Alg          jwt.SigningMethod
	SignKey      any
	VerifyKey    any
	ParseFn      ParseFn
	ValidatorFns []ValidatorFn
}

func New(secretKey any, opts ...Option) *JwtAuth {
	ja := &JwtAuth{
		Alg:          DefaultAlg,
		SignKey:      secretKey,
		VerifyKey:    secretKey,
		ParseFn:      DefaultParseFn,
		ValidatorFns: make([]ValidatorFn, 0),
	}

	for _, opt := range opts {
		opt(ja)
	}

	return ja
}
