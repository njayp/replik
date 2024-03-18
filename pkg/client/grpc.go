package client

import (
	"context"

	"github.com/njayp/replik/pkg/api"
	"github.com/njayp/replik/pkg/conn"
	"github.com/njayp/replik/pkg/manager"
)

type Client struct {
	manager Manager
}

var DefaultClient = Client{manager: manager.NewManager()}

func (c *Client) GetFile(ctx context.Context, filename string) error {
	client := conn.NewClient()
	stream, err := client.File(ctx, &api.FileRequest{Filename: filename})
	if err != nil {
		return err
	}

	ch := make(chan *api.Chunk)
	go c.manager.WriteFileFromCh(ctx, filename, ch)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			chunk, err := stream.Recv()
			if err != nil {
				return err
			}
			ch <- chunk
		}

	}
}
