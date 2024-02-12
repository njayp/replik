package manager

import (
	"context"
	"testing"
	"time"

	"github.com/njayp/replik/pkg/api"
)

func TestRW(t *testing.T) {
	ctx := context.Background()
	filename := "coon.jpg"
	manager := NewManager()
	ch := manager.ReadPart(ctx, &api.FileRequest{Filename: filename})
	go manager.WritePart(ctx, filename, ch)
	time.Sleep(time.Second * 2)
}
