package server

import (
	"context"

	"github.com/njayp/replik/pkg/api"
	"github.com/njayp/replik/pkg/conn"
	"github.com/njayp/replik/pkg/manager"
	"google.golang.org/grpc"
)

type Service struct {
	api.UnimplementedReplikServer
	manager Manager
}

func NewService() error {
	lis := conn.NewListener()
	s := grpc.NewServer()
	// manager chosen here
	api.RegisterReplikServer(s, &Service{manager: manager.NewManager()})
	return s.Serve(lis)
}

func (s *Service) File(req *api.FileRequest, stream api.Replik_FileServer) error {
	ctx := stream.Context()
	ch := s.manager.ReadFileToCh(ctx, req)
	for {
		select {
		case <-ctx.Done():
			return nil
		case chunk := <-ch:
			stream.Send(chunk)
		}

		// if ch is nil, we are done reading
		if ch == nil {
			break
		}
	}

	return nil
}

func (s *Service) Status(context.Context, *api.Empty) (*api.StatusResponse, error) {
	return &api.StatusResponse{Status: "alive"}, nil
}
