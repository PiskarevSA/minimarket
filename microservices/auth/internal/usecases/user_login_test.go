package usecases

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/suite"

	"github.com/PiskarevSA/minimarket/microservices/auth/internal/usecases/mocks"
)

type userLogInSuite struct {
	suite.Suite
	mockStorage *mocks.MockUserStorage
	usecase     *UserLogIn
}

func TestUserLogIn(t *testing.T) {
	suite.Run(t, new(userLogInSuite))
}

func (s *userLogInSuite) SetupTest() {
	s.mockStorage = new(mocks.MockUserStorage)
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
