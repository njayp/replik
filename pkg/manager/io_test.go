package manager

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/njayp/replik/pkg/api"
)

var folder = "output"
var fileName = "test.txt"
var path = filepath.Join(folder, fileName)
var file2Name = "test2.txt"
var path2 = filepath.Join(folder, file2Name)

func TestR(t *testing.T) {
	// Create a string
	s := "Hello, world!"

	// Create a buffer from the string
	buffer := bytes.NewBufferString(s)

	// Create a buffered writer from the buffer
	r := bufio.NewReader(buffer)
	ch := make(chan *api.Chunk)
	go ReadToCh(context.Background(), r, ch)

	chunk := <-ch
	bytes := chunk.Data[:chunk.Size]
	data := string(bytes)
	if data != s {
		t.Error(data)
	}
}

func TestW(t *testing.T) {
	ch := make(chan *api.Chunk)

	os.MkdirAll(folder, os.ModePerm)
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	// defer are FILO, flush then close
	defer w.Flush()

	go WriteFromCh(context.Background(), w, ch)
	for i := 1; i < 999; i++ {
		line := fmt.Sprintf("this is line %v\n", i)
		bytes := []byte(line)
		size := len(bytes)
		ch <- &api.Chunk{Data: bytes, Size: int64(size)}
	}
	time.Sleep(time.Second)
}

func TestRW(t *testing.T) {
	ctx := context.Background()
	inf, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	outf, err := os.Create(path2)
	if err != nil {
		panic(err)
	}

	ch := make(chan *api.Chunk)
	go ReadToCh(ctx, bufio.NewReader(inf), ch)
	outbuf := bufio.NewWriter(outf)
	go WriteFromCh(ctx, outbuf, ch)

	time.Sleep(time.Second)
	// flush last chunk into file
	outbuf.Flush()

}
