package client

import (
	"context"
	"sync"

	"github.com/njayp/replik/pkg/api"
)

func Get(ctx context.Context, path string) {
	client := NewClient()

	files, err := client.List(ctx, path)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	for _, file := range files.Files {
		wg.Add(1)

		go func(file *api.File) {
			defer wg.Done()
			err := client.File(ctx, file.Path)
			if err != nil {
				println(err)
			}
		}(file)
	}
	wg.Wait()
}
