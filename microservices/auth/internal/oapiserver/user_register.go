package oapiserver

import (
	"context"

	"github.com/PiskarevSA/minimarket/microservices/auth/internal/oapiserver/oapi"
)

func (s *OapiServer) UserRegister(ctx context.Context, request oapi.UserRegisterRequestObject) (
	oapi.UserRegisterResponseObject,
	error,
) {
	// request.
	// passwordSalt := strgen.Global.Generate(24)
	return nil, nil
}
