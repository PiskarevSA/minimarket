package usecases

import (
	"testing"
	"time"

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
		time.Hour,
		12*time.Hour,
	)
}

func (s *userLogInSuite) TestUserLogIn() {
}
