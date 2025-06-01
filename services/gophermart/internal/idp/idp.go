package idp

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type IdentityProvider struct {
	issuer        string
	signingMethod jwt.SigningMethod
	signingKey    []byte
	tokenTTL      time.Duration
}

func NewIdentityProvider(
	issuer string,
	signingMethod jwt.SigningMethod,
	signingKey []byte,
	tokenTTL time.Duration,
) *IdentityProvider {
	return &IdentityProvider{
		issuer:        issuer,
		signingMethod: signingMethod,
		signingKey:    signingKey,
		tokenTTL:      tokenTTL,
	}
}

func (i *IdentityProvider) IssueToken(
	userID uuid.UUID,
	now time.Time,
) (tokenString string, err error) {
	iat := now.UTC().Unix()
	exp := now.Add(i.tokenTTL).Unix()

	token := jwt.NewWithClaims(
		i.signingMethod,
		jwt.MapClaims{
			"iss": i.issuer,
			"sub": userID.String(),
			"iat": iat,
			"exp": exp,
		},
	)

	return token.SignedString(i.signingKey)
}
