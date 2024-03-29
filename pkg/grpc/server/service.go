package server

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/fs"
	"net"
	"os"
	"path/filepath"

	"github.com/njayp/gcm"
	"github.com/njayp/replik/pkg/api"
	"github.com/njayp/replik/pkg/config"
	"google.golang.org/grpc"
)

type Service struct {
	api.UnimplementedReplikServer
}

func NewService() error {
	env := gcm.Env[config.Env]()
	url := fmt.Sprintf("%s:%v", env.Address, env.Port)
	lis, err := net.Listen("tcp", url)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(
		gcm.NewGcmServerOpts(gcm.WithEnv[config.Env])...,
	)
	api.RegisterReplikServer(s, &Service{})
	return s.Serve(lis)
}

func (s *Service) GetFileList(ctx context.Context, r *api.FileListRequest) (*api.FileList, error) {
	var files []*api.File

	fileInfo, err := os.Stat(r.Path)
	if err != nil {
		return nil, err
	}

	if !fileInfo.IsDir() {
		files = append(files, &api.File{Path: r.Path})
		return &api.FileList{Files: files}, nil
	}

	err = filepath.WalkDir(r.Path, func(path string, d fs.DirEntry, err error) error {
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

func (s *Service) GetFile(r *api.FileRequest, stream api.Replik_GetFileServer) error {
	ctx := stream.Context()
	file, err := os.Open(r.Path)
	if err != nil {
		return err
	}
	defer file.Close()
	buf := bufio.NewReader(file)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			bytes := make([]byte, gcm.GetEnv[config.Env](ctx).ChunkSize)
			n, err := buf.Read(bytes)
			if err != nil && err != io.EOF {
				return err
			}

			if n == 0 {
				return nil
			}

			stream.Send(&api.Chunk{Data: bytes, Size: int64(n)})
		}
	}
}

func (s *Service) GetStatus(context.Context, *api.Empty) (*api.Status, error) {
	return &api.Status{Status: "ok"}, nil
}
