package oapiserver

import (
	"context"
	"errors"

	"github.com/PiskarevSA/minimarket/microservices/auth/internal/oapiserver/oapi"
	"github.com/PiskarevSA/minimarket/microservices/auth/internal/usecases"
)

func userRegisterHandleUsecaseError(
	err error,
) oapi.UserRegisterResponseObject {
	if errors.Is(err, usecases.ErrLoginAlreadyInUse) {
		return oapi.UserRegister409JSONResponse{
			Code:    oapi.U1956,
			Message: oapi.LoginAlreadyInUse,
		}
	}

	return oapi.UserRegister500JSONResponse{
		Code:    oapi.S1394,
		Message: oapi.InternalServerError,
	}
}

func (s *OapiServer) UserRegister(
	ctx context.Context,
	request oapi.UserRegisterRequestObject,
) (
	oapi.UserRegisterResponseObject,
	error,
) {
	userRegistered, err := s.userRegister.Do(
		ctx,
		request.Body.Login,
		request.Body.Password,
	)
	if err != nil {
		return userRegisterHandleUsecaseError(err), nil
	}

	return oapi.UserRegister200JSONResponse{
		Id:           userRegistered.UserId,
		Login:        userRegistered.Login,
		AccessToken:  userRegistered.AccessToken,
		RefreshToken: userRegistered.RefreshToken,
	}, nil
}
