package manager

import (
	"bufio"
	"context"
	"io"
	"path/filepath"

	"github.com/njayp/replik/pkg/api"
)

var chunkSize = 4096

// file is read into ch async
func (m *Manager) ReadFileToCh(ctx context.Context, req *api.FileRequest) <-chan *api.Chunk {
	// TODO rm
	path := filepath.Join("assets", req.Filename)
	file, err := m.EnsureFileRO(path)
	if err != nil {
		// TODO handle err
		panic(err)
	}

	ch := make(chan *api.Chunk)
	go ReadToCh(ctx, bufio.NewReader(file), ch)
	return ch
}

// reads from r into ch
func ReadToCh(ctx context.Context, r io.Reader, ch chan<- *api.Chunk) error {
	for {
		// make new buffer for each chunk
		buf := make([]byte, chunkSize)
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			// TODO handle err
			return err
		}
		// if EOF and buffer is empty, set ch to nil to signal end of data
		if n == 0 {
			println("coonable")
			ch = nil
			return nil
		}

		select {
		case <-ctx.Done():
			return nil // dead stream
		case ch <- &api.Chunk{Data: buf, Size: int64(n)}:
		}
	}
}
