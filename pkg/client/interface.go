package client

import (
	"context"

	"github.com/njayp/replik/pkg/api"
)

type Manager interface {
	WritePart(ctx context.Context, filename string, ch <-chan *api.Chunk) error
}
