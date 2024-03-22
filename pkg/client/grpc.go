package client

import (
	"context"
	"io"

	"github.com/njayp/replik/pkg/api"
	"github.com/njayp/replik/pkg/conn"
	"github.com/njayp/replik/pkg/manager"
)

type Client struct {
	manager *manager.Manager
}

var DefaultClient = Client{manager: manager.NewManager()}

func (c *Client) GetFile(ctx context.Context, path string) error {
	client := conn.NewClient()
	stream, err := client.GetFile(ctx, &api.FileRequest{Path: path})
	if err != nil {
		return err
	}

	ch := make(chan *api.Chunk)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go c.manager.WriteFileFromCh(ctx, path, ch)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			chunk, err := stream.Recv()
			if err != nil {
				// graceful end of stream
				if err == io.EOF {
					return nil
				}
				return err
			}
			ch <- chunk
		}
	}
}
