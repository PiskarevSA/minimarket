package usecases

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/suite"
)

type userRegisterSuite struct {
	suite.Suite
	mockStorage *MockUserStorage
	usecase     *UserRegister
}

func TestUserRegister(t *testing.T) {
	suite.Run(t, new(userLogInSuite))
}

func (s *userRegisterSuite) SetupTest() {
	s.mockStorage = new(MockUserStorage)
	s.usecase = NewUserRegister(
		s.mockStorage,
		[]byte("jwt"),
		jwt.SigningMethodHS256,
		time.Hour,
		12*time.Hour,
	)
}

func (s *userRegisterSuite) TestUserLogIn() {

}
