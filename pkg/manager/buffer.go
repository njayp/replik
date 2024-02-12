package manager

import (
	"os"
)

func NewManager() *Manager {
	return &Manager{buffer: make(map[string]*os.File)}
}

type Manager struct {
	buffer map[string]*os.File
}

// TODO combine funcs
func (fm *Manager) EnsureFileRO(path string) (*os.File, error) {
	cache, ok := fm.buffer[path]
	if ok {
		return cache, nil
	}

	// file not in cache
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	fm.buffer[path] = file
	return file, nil
}

func (b *Manager) close() {
	for _, v := range b.buffer {
		v.Close()
	}
}
