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
	status, err := client.GetStatus(ctx, &api.Empty{})
	if err != nil {
		t.Fatal(err)
	}

	if status.Status != "alive" {
		t.Error(status.Status)
	}
}

func TestGetFile(t *testing.T) {
	DefaultClient.GetFile(context.Background(), "coon.jpg")
}
