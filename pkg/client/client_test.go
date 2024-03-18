package client

import (
	"context"
	"testing"

	"github.com/njayp/replik/pkg/api"
	"github.com/njayp/replik/pkg/conn"
)

func TestStatus(t *testing.T) {
	ctx := context.Background()
	client := conn.NewClient()
	status, err := client.Status(ctx, &api.Empty{})
	if err != nil {
		t.Error(err)
		return
	}

	if status.Status != "alive" {
		t.Error(status.Status)
	}
}

func TestGetFile(t *testing.T) {
	t.Error(DefaultClient.GetFile(context.Background(), "coon.jpg"))
}
