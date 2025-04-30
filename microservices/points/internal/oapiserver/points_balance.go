package oapiserver

import (
	"context"

	"github.com/PiskarevSA/minimarket/microservices/points/internal/oapiserver/oapi"
)

func (s *OapiServer) PointsBalance(ctx context.Context, _ oapi.PointsBalanceRequestObject) (oapi.PointsBalanceResponseObject, error) {
	return oapi.PointsBalance401JSONResponse{
		Code:    oapi.U1527,
		Message: oapi.UserUnauthenticated,
	}, nil
}
