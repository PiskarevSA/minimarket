package usecases

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"
	"time"

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

func (s *userRegisterSuite) TestUserLogInSuccess() {
	s.mockStorage.CreateUserFunc = func(
		ctx context.Context, userId uuid.UUID, login string, passwordHash string,
	) error {
		return nil
	}

	s.generatePrivateKey()

	user, err := s.usecase.Do(context.Background(), "login", "password")
	require.NoError(s.T(), err)

	require.NotEmpty(s.T(), user.UserId)
	require.Equal(s.T(), "login", user.Login)
	require.NotEmpty(s.T(), user.AccessToken)
	require.NotEmpty(s.T(), user.RefreshToken)
}

func (s *userRegisterSuite) generatePrivateKey() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(s.T(), err)

	s.usecase.jwtSignKey = privateKey
}
