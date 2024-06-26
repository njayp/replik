package client

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/njayp/replik/pkg/api"
	"github.com/njayp/replik/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client api.ReplikClient
}

func NewClient() *Client {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	url := fmt.Sprintf("%s:%v", config.Env.Address, config.Env.Port)
	conn, err := grpc.Dial(url, opts...)
	if err != nil {
		panic(err)
	}
	return &Client{client: api.NewReplikClient(conn)}
}

func (c *Client) List(ctx context.Context, path string) (*api.FileList, error) {
	return c.client.GetFileList(ctx, &api.FileListRequest{Path: path})
}

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func (c *Client) File(ctx context.Context, path string) error {
	stream, err := c.client.GetFile(ctx, &api.FileRequest{Path: path})
	if err != nil {
		return err
	}
	file, err := create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	buf := bufio.NewWriter(file)
	defer buf.Flush() // defer is FILO

	for {
		select {
		case <-stream.Context().Done():
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
			buf.Write(chunk.GetData()[:chunk.GetSize()])
		}
	}
}

func (c *Client) Status(ctx context.Context) (*api.Status, error) {
	return c.client.GetStatus(ctx, &api.Empty{})
}
