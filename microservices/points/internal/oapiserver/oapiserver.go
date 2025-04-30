package oapiserver

import (
	"github.com/PiskarevSA/minimarket/microservices/points/internal/oapiserver/oapi"
)

var _ oapi.StrictServerInterface = (*OapiServer)(nil)

type OapiServer struct{}

func New() *OapiServer {
	return nil
}
