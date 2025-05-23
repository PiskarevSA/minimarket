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
	tokenTtl      time.Duration
}

func NewIdentityProvider(
	issuer string,
	signingMethod jwt.SigningMethod,
	signingKey []byte,
	tokenTtl time.Duration,
) *IdentityProvider {
	return &IdentityProvider{
		issuer:        issuer,
		signingMethod: signingMethod,
		signingKey:    signingKey,
		tokenTtl:      tokenTtl,
	}
}

func (i *IdentityProvider) IssueToken(
	userId uuid.UUID,
	now time.Time,
) (tokenString string, err error) {
	iat := now.UTC().Unix()
	exp := now.Add(i.tokenTtl).Unix()

	token := jwt.NewWithClaims(
		i.signingMethod,
		jwt.MapClaims{
			"iss": i.issuer,
			"sub": userId.String(),
			"iat": iat,
			"exp": exp,
		},
	)

	return token.SignedString(i.signingKey)
}
