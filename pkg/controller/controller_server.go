package controller

import (
	api "hsbc_connector/pkg/api"
	sidecar "hsbc_connector/sidecar"
)

import (
	"context"
)

type Server struct {
	sidecar.UnimplementedServiceServer
}

func (s *Server) Sidecar(ctx context.Context, request *sidecar.SidecarRequest) (*sidecar.ResponseDto, error) {

	result, err := api.Sidecar(request.RequestDto, ctx)

	return result, err
}
