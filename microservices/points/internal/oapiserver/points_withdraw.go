package oapiserver

import (
	"context"

	"github.com/PiskarevSA/minimarket/microservices/points/internal/oapiserver/oapi"
)

func (s *OapiServer) PointsWithdraw(ctx context.Context, request oapi.PointsWithdrawRequestObject) (oapi.PointsWithdrawResponseObject, error) {
	return oapi.PointsWithdraw422JSONResponse{
		Code:    oapi.U1062,
		Message: oapi.InvalidOrderNumber,
	}, nil
}
