package server

import (
	"context"
	"testing"

	"github.com/njayp/replik/pkg/api"
	"github.com/njayp/replik/pkg/conn"
)

func TestServer(t *testing.T) {
	ctx := context.Background()
	client := conn.NewClient()
	status, err := client.Status(ctx, &api.Empty{})
	if err != nil {
		t.Error(err)
		return
	}

	t.Error(status.Status)
}
