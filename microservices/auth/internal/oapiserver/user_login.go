package oapiserver

import (
	"context"

	"github.com/PiskarevSA/minimarket/microservices/auth/internal/oapiserver/oapi"
)

func (s *OapiServer) UserLogin(ctx context.Context, request oapi.UserLoginRequestObject) (
	oapi.UserLoginResponseObject,
	error,
) {
	return nil, nil
}
