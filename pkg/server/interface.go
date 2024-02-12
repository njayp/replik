package server

import (
	"context"

	"github.com/njayp/replik/pkg/api"
)

type Manager interface {
	ReadPart(ctx context.Context, req *api.FileRequest) <-chan *api.Chunk
}
