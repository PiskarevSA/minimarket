package usecases

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"testing"
	"time"

	"github.com/PiskarevSA/minimarket/microservices/auth/internal/storage"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

type userLogInSuite struct {
	suite.Suite
	mockStorage *MockUserStorage
	usecase     *UserLogIn
}

func TestUserLogIn(t *testing.T) {
	suite.Run(t, new(userLogInSuite))
}

func (s *userLogInSuite) SetupTest() {
	s.mockStorage = new(MockUserStorage)
	s.usecase = NewUserLogIn(
		s.mockStorage,
		[]byte(""),
		time.Hour,
		12*time.Hour,
	)
}

func (s *userLogInSuite) TestUserLogInSuccess() {
	s.mockStorage.GetUserCredsFunc = func(ctx context.Context, login string,
	) (*storage.UserCreds, error) {
		return &storage.UserCreds{
			UserId:       uuid.New(),
			PasswordHash: s.generatePasswordHash("password"),
		}, nil
	}

	s.generatePrivateKey()

	user, err := s.usecase.Do(context.Background(), "login", "password")
	require.NoError(s.T(), err)

	s.validateUuid(user.UserId)
	require.Equal(s.T(), "login", user.Login)
	s.validateToken(user.AccessToken)
	s.validateToken(user.RefreshToken)
}

func (s *userLogInSuite) generatePrivateKey() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(s.T(), err)

	s.usecase.jwtSignKey = privateKey
}

func (s *userLogInSuite) generatePasswordHash(password string) string {
	passwordBytes, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost)
	require.NoError(s.T(), err)
	return string(passwordBytes)
}

func (s *userLogInSuite) validateUuid(uuisString string) {
	userUuid, err := uuid.Parse(uuisString)
	require.NoError(s.T(), err)
	require.Equal(s.T(), uuid.Version(0x04), userUuid.Version())
}

func (s *userLogInSuite) validateToken(tokenString string) {
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
