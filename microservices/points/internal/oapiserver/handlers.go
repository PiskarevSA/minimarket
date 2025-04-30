package oapiserver

import (
	"context"

	"github.com/PiskarevSA/minimarket/microservices/points/internal/oapiserver/oapi"
)

var _ oapi.StrictServerInterface = (*Server)(nil)

type Server struct{}

func New() *Server {
	return nil
}

func (h *Server) PointsBalance(ctx context.Context, _ oapi.PointsBalanceRequestObject) (oapi.PointsBalanceResponseObject, error) {
	return oapi.PointsBalance401JSONResponse{
		Code:    oapi.U1527,
		Message: oapi.UserUnauthenticated,
	}, nil
}

func (h *Server) PointsWithdraw(ctx context.Context, request oapi.PointsWithdrawRequestObject) (oapi.PointsWithdrawResponseObject, error) {
	return oapi.PointsWithdraw422JSONResponse{
		Code:    oapi.U1062,
		Message: oapi.InvalidOrderNumber,
	}, nil
}
