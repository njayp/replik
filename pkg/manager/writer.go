package manager

import (
	"bufio"
	"context"
	"io"
	"os"

	"github.com/njayp/replik/pkg/api"
)

// blocks. cancel ctx to flush and close
func (m *Manager) WriteFileFromCh(ctx context.Context, path string, ch <-chan *api.Chunk) error {
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
