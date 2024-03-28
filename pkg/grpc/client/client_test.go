package client

import (
	"context"
	"testing"
)

func TestStatus(t *testing.T) {
	status, err := NewClient().Status(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if status.Status != "ok" {
		t.Error(status.Status)
	}
}

func TestGetFile(t *testing.T) {
	err := NewClient().File(context.Background(), "test/coon.jpg")
	if err != nil {
		t.Error(err)
	}
}

func TestGet(t *testing.T) {
	Get(context.Background(), "test")
}
