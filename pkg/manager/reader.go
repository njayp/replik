package manager

import (
	"bufio"
	"context"
	"io"
	"os"

	"github.com/njayp/replik/pkg/api"
)

const chunkSize = 64 * 1024 // 64 KiB

// file is read into ch async, ch is closed when done
func (m *Manager) ReadFileToCh(ctx context.Context, req *api.FileRequest) <-chan *api.Chunk {
	file, err := m.EnsureFile(req.Path, os.Open)
	if err != nil {
		// TODO handle err
		panic(err)
	}

	// set file to requested start index
	file.Seek(req.Index, io.SeekStart)
	ch := make(chan *api.Chunk)
	go ReadToCh(ctx, bufio.NewReader(file), ch)
	return ch
}

// reads from r into ch, ch is closed when done
func ReadToCh(ctx context.Context, r io.Reader, ch chan<- *api.Chunk) error {
	for {
		// make new buffer for each chunk
		buf := make([]byte, chunkSize)
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			// close ch to signal EOF
			close(ch)
			return nil
		}

		select {
		case <-ctx.Done():
			return nil // dead stream
		case ch <- &api.Chunk{Data: buf, Size: int64(n)}:
		}
	}
}
