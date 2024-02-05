package file

import (
	"context"
	"math"
	"os"
)

const chunkSize = 64
const maxBufferSize = 10

func ReadFile(ctx context.Context, path string) <-chan []byte {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	stats, err := f.Stat()
	if err != nil {
		panic(err)
	}

	cap := buffSize(int(stats.Size()))
	ch := make(chan []byte, cap)

	// create a thread to read a file and send chunks as they are ready
	go func() {
		defer f.Close()
		for {
			chunk := make([]byte, chunkSize)
			// TODO make sure noBytes is correct
			noBytes, err := f.Read(chunk)
			if err != nil {
				panic(err)
			}

			if noBytes == 0 {
				return // nothing left to read
			}

			select {
			case <-ctx.Done():
				return // dead stream
			default:
				ch <- chunk
			}
		}
	}()

	return ch
}

func noChunks(size int) int {
	return int(math.Ceil(float64(size) / float64(chunkSize)))
}

func buffSize(size int) int {
	noChunks := noChunks(size)
	if noChunks > maxBufferSize {
		return maxBufferSize
	}
	return noChunks
}
