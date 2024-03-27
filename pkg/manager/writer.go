package manager

import (
	"bufio"
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/njayp/replik/pkg/api"
)

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

// blocks. cancel ctx to flush
func (m *Manager) WriteFileFromCh(ctx context.Context, path string, ch <-chan *api.Chunk) error {
	f, err := m.EnsureFile(path, create)
	if err != nil {
		return nil
	}

	buf := bufio.NewWriter(f)
	defer buf.Flush()
	return WriteFromCh(ctx, buf, ch)
}

// blocks. writes bytes from channel until ctx is cancelled
func WriteFromCh(ctx context.Context, w io.Writer, ch <-chan *api.Chunk) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case chunk := <-ch:
			// TODO look at again
			if _, err := w.Write(chunk.GetData()[:chunk.GetSize()]); err != nil {
				return err
			}
		}
	}
}
