package usecases

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createAccessToken(
	userId string,
	expiry time.Duration,
	signKey any,
	algo jwt.SigningMethod,
) (string, error) {
	now := time.Now()

	iat := now.UTC().Unix()
	exp := now.Add(time.Hour * time.Duration(expiry)).Unix()

	token := jwt.NewWithClaims(algo,
		jwt.MapClaims{
			"iss": "auth-service",
			"sub": userId,
			"exp": exp,
			"iat": iat,
		},
	)

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", fmt.Errorf(
			"failed to create access token: %w",
			err,
		)
	}

	return tokenString, err
}

func createRefreshToken(
	userId string,
	expiry time.Duration,
	signKey any,
	algo jwt.SigningMethod,
) (string, error) {
	now := time.Now()

	iat := now.UTC().Unix()
	exp := now.Add(time.Hour * time.Duration(expiry)).Unix()

	token := jwt.NewWithClaims(algo,
		jwt.MapClaims{
			"iss": "auth-service",
			"sub": userId,
			"exp": exp,
			"iat": iat,
		},
	)

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", fmt.Errorf(
			"failed to create refresh token: %w",
			err,
		)
	}

	return tokenString, err
}

func createTokenPair(
	userId string,
	accessTokenExpiry time.Duration,
	refreshTokenExpiry time.Duration,
	signKey any,
	algo jwt.SigningMethod,
) (string, string, error) {
	accessToken, err := createAccessToken(
		userId,
		accessTokenExpiry,
		signKey,
		algo,
	)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := createRefreshToken(
		userId,
		refreshTokenExpiry,
		signKey,
		algo,
	)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
