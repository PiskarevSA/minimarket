package usecases

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/suite"
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
		[]byte("jwt"),
		jwt.SigningMethodHS256,
		time.Hour,
		12*time.Hour,
	)
}

func (s *userLogInSuite) TestUserLogIn() {

}
