package server

import (
	"context"
	"encoding/json"

	"github.com/njayp/replik/pkg/api"
	"github.com/njayp/replik/pkg/conn"
	"github.com/njayp/replik/pkg/manager"
	"google.golang.org/grpc"
)

type Service struct {
	api.UnimplementedReplikServer
	manager *manager.Manager
}

func NewService() error {
	lis := conn.NewListener()
	s := grpc.NewServer()
	api.RegisterReplikServer(s, &Service{manager: manager.NewManager()})
	return s.Serve(lis)
}

func (s *Service) GetPathInfo(context.Context, *api.PathInfoRequest) (*api.PathInfo, error) {
	// TODO add tree stuct getter to manager
	tree, err := json.Marshal("testo")
	if err != nil {
		return nil, err
	}
	return &api.PathInfo{Tree: tree}, nil
}

func (s *Service) GetFile(req *api.FileRequest, stream api.Replik_GetFileServer) error {
	ctx := stream.Context()
	ch := s.manager.ReadFileToCh(ctx, req)
	for {
		select {
		case <-ctx.Done():
			return nil
		case chunk := <-ch:
			// ch is closed
			if chunk == nil {
				return nil
			}
			stream.Send(chunk)
		}
	}
}

func (s *Service) GetStatus(context.Context, *api.Empty) (*api.Status, error) {
	return &api.Status{Status: "alive"}, nil
}
