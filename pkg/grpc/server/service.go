package server

import (
	"context"
	"io/fs"
	"path/filepath"

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
	m := manager.NewManager()
	defer m.Close()
	api.RegisterReplikServer(s, &Service{manager: m})
	return s.Serve(lis)
}

func (s *Service) GetFileList(ctx context.Context, r *api.FileListRequest) (*api.FileList, error) {
	var files []*api.File
	err := filepath.WalkDir(r.Path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, &api.File{Path: path})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &api.FileList{Files: files}, nil
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
	return &api.Status{Status: "ok"}, nil
}
