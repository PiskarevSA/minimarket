package oapiserver

import (
	"context"
	"errors"

	"github.com/PiskarevSA/minimarket/microservices/auth/internal/oapiserver/oapi"
	"github.com/PiskarevSA/minimarket/microservices/auth/internal/usecases"
)

func userLogInHandleUsecaseError(
	err error,
) oapi.UserLogInResponseObject {
	if errors.Is(err, usecases.ErrInvalidLoginOrPassword) {
		return oapi.UserLogIn401JSONResponse{
			Code:    oapi.U1385,
			Message: oapi.InvalidLoginOrPassword,
		}
	}

	return oapi.UserLogIn500JSONResponse{
		Code:    oapi.S1394,
		Message: oapi.InternalServerError,
	}
}

func (s *OapiServer) UserLogIn(
	ctx context.Context,
	request oapi.UserLogInRequestObject,
) (
	oapi.UserLogInResponseObject,
	error,
) {
	userLoggedIn, err := s.userLogIn.Do(
		ctx,
		request.Body.Login,
		request.Body.Password,
	)
	if err != nil {
		return userLogInHandleUsecaseError(err), nil
	}

	return oapi.UserLogIn200JSONResponse{
		Id:           userLoggedIn.UserId,
		Login:        userLoggedIn.Login,
		AccessToken:  userLoggedIn.AccessToken,
		RefreshToken: userLoggedIn.RefreshToken,
	}, nil
}
