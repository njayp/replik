package main

import (
	"context"

	"github.com/njayp/replik/pkg/client"
)

func main() {
	client.DefaultClient.GetFile(context.Background(), "coon.jpg")
}
