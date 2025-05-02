package oapiserver

import (
	"context"

	"github.com/PiskarevSA/minimarket/microservices/auth/internal/models"
	"github.com/PiskarevSA/minimarket/microservices/auth/internal/oapiserver/oapi"
)

type (
	UserRegisterUsecase interface {
		Do(
			ctx context.Context,
			login, password string,
		) (*models.UserRegistered, error)
	}

	UserLogInUsecase interface {
		Do(
			ctx context.Context,
			login, password string,
		) (*models.UserLoggedIn, error)
	}
)

var _ oapi.StrictServerInterface = (*OapiServer)(nil)

type OapiServer struct {
	userRegister UserRegisterUsecase
	userLogIn    UserLogInUsecase
}

func New(
	userRegister UserRegisterUsecase,
	userLogIn UserLogInUsecase,
) *OapiServer {
	return &OapiServer{
		userRegister: userRegister,
		userLogIn:    userLogIn,
	}
}
