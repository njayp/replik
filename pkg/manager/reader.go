package manager

import (
	"bufio"
	"context"
	"io"
	"path/filepath"

	"github.com/njayp/replik/pkg/api"
)

var chunkSize = 4096

func (fm *Manager) ReadPart(ctx context.Context, req *api.FileRequest) <-chan *api.Chunk {
	path := filepath.Join("assets", req.Filename)
	file, err := fm.EnsureFileRO(path)
	if err != nil {
		// TODO handle err
		panic(err)
	}

	ch := make(chan *api.Chunk)
	go func() {
		defer close(ch)
		r := bufio.NewReader(file)
		buf := make([]byte, chunkSize)

		for {
			// read chunk
			n, err := r.Read(buf)
			println("blarg", n)
			if err != nil && err != io.EOF {
				// TODO handle err
				panic(err)
			}
			if n == 0 {
				return
			}

			// send data if stream is still active
			select {
			case <-ctx.Done():
				return // dead stream
			default:
				ch <- &api.Chunk{Data: buf, Size: int64(n)}
			}
		}
	}()

	return ch
}
