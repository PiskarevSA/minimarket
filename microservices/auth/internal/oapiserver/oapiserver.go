package oapiserver

import (
	"context"

	"github.com/PiskarevSA/minimarket/microservices/auth/internal/oapiserver/oapi"
)

type UserUsecase interface {
	Register(ctx context.Context, login, password string) (accessToken string, err error)
	Login(ctx context.Context, login, password string) (accessToken string, err error)
}

var _ oapi.StrictServerInterface = (*OapiServer)(nil)

type OapiServer struct{ usecase UserUsecase }

func New(usecase UserUsecase) *OapiServer {
	return &OapiServer{usecase}
}
