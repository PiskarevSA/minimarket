package usecases

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type userRegisterSuite struct {
	suite.Suite
	mockStorage *MockUserStorage
	usecase     *UserRegister
}

func TestUserRegister(t *testing.T) {
	suite.Run(t, new(userRegisterSuite))
}

func (s *userRegisterSuite) SetupTest() {
	s.mockStorage = new(MockUserStorage)
	s.usecase = NewUserRegister(
		s.mockStorage,
		[]byte(""),
		time.Hour,
		12*time.Hour,
	)
}

func (s *userRegisterSuite) TestUserRegisterSuccess() {
	s.mockStorage.CreateUserFunc = func(
		ctx context.Context, userId uuid.UUID, login string, passwordHash string,
	) error {
		return nil
	}

	s.generatePrivateKey()

	user, err := s.usecase.Do(context.Background(), "login", "password")
	require.NoError(s.T(), err)

	s.validateUuid(user.UserId)
	require.Equal(s.T(), "login", user.Login)
	s.validateToken(user.AccessToken)
	s.validateToken(user.RefreshToken)
}

func (s *userRegisterSuite) generatePrivateKey() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(s.T(), err)

	s.usecase.jwtSignKey = privateKey
}

func (s *userRegisterSuite) validateUuid(uuisString string) {
	userUuid, err := uuid.Parse(uuisString)
	require.NoError(s.T(), err)
	require.Equal(s.T(), uuid.Version(0x04), userUuid.Version())
}

func (s *userRegisterSuite) validateToken(tokenString string) {
	keyFunc := func(t *jwt.Token) (any, error) {
		if t.Method.Alg() != "ES256" {
			return nil, fmt.Errorf("wrong algo")
		}

		// jwt.Parse invokes ECDSA parser, that expect *ecdsa.PublicKey
		return &s.usecase.jwtSignKey.(*ecdsa.PrivateKey).PublicKey, nil
	}

	token, err := jwt.Parse(tokenString, keyFunc)
	require.NoError(s.T(), err)
	require.True(s.T(), token.Valid)
}
