package jwtauth

import "github.com/golang-jwt/jwt/v5"

type Option func(*JwtAuth)

func WithAlg(alg jwt.SigningMethod) Option {
	return func(a *JwtAuth) {
		a.Alg = alg
	}
}

func WithSignKey(key any) Option {
	return func(a *JwtAuth) {
		a.SignKey = key
	}
}

func WithVerifyKey(key any) Option {
	return func(a *JwtAuth) {
		a.VerifyKey = key
	}
}

func WithValidator(fn ValidatorFn) Option {
	return func(a *JwtAuth) {
		a.ValidatorFns = append(a.ValidatorFns, fn)
	}
}
