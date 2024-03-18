package manager

import (
	"bufio"
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/njayp/replik/pkg/api"
)

// blocks until ctx is cancelled
func (m *Manager) WriteFileFromCh(ctx context.Context, filename string, ch <-chan *api.Chunk) error {
	path := filepath.Join("output", filename)
	// TODO use manager
	f, err := os.Create(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	buf := bufio.NewWriter(f)
	defer buf.Flush() // defer is FILO
	return WriteFromCh(ctx, buf, ch)
}

// blocks. uses writer to write bytes from channel until ctx is cancelled
func WriteFromCh(ctx context.Context, w io.Writer, ch <-chan *api.Chunk) error {
	for {
		select {
		case <-ctx.Done():
			println("writer ctx cancelled")
			return nil
		case chunk := <-ch:
			if _, err := w.Write(chunk.GetData()[:chunk.GetSize()]); err != nil {
				return err
			}
		}
	}
}
