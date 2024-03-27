package client

import (
	"context"
	"io"

	"github.com/njayp/replik/pkg/api"
	"github.com/njayp/replik/pkg/conn"
	"github.com/njayp/replik/pkg/manager"
)

type Client struct {
	client  api.ReplikClient
	manager *manager.Manager
}

func NewClient() *Client {
	return &Client{client: api.NewReplikClient(conn.NewConn()), manager: manager.NewManager()}
}

func (c *Client) List(ctx context.Context, path string) (*api.FileList, error) {
	return c.client.GetFileList(ctx, &api.FileListRequest{Path: path})
}

func (c *Client) File(ctx context.Context, path string) error {
	stream, err := c.client.GetFile(ctx, &api.FileRequest{Path: path})
	if err != nil {
		return err
	}
	ctx = stream.Context()
	ch := make(chan *api.Chunk)
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

func (c *Client) Status(ctx context.Context) (*api.Status, error) {
	return c.client.GetStatus(ctx, &api.Empty{})
}
