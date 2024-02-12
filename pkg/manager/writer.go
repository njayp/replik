package manager

import (
	"bufio"
	"context"
	"os"
	"path/filepath"

	"github.com/njayp/replik/pkg/api"
)

// TODO seperate from reader
func (fm *Manager) WritePart(ctx context.Context, filename string, ch <-chan *api.Chunk) error {
	dir := "output"
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	path := filepath.Join(dir, filename)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// defer are f-in-l-out, flush then close
	w := bufio.NewWriter(f)
	defer w.Flush()

	for {
		select {
		case <-ctx.Done():
			return nil
		case chunk := <-ch:
			if chunk.GetSize() == 0 {
				return nil // nil chunk, closed channel
			}
			if _, err := w.Write(chunk.GetData()[:chunk.GetSize()]); err != nil {
				panic(err)
			}
		}
	}
}
