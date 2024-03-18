package manager

import (
	"os"
)

func NewManager() *Manager {
	return &Manager{files: make(map[string]*os.File)}
}

type Manager struct {
	files map[string]*os.File
}

// TODO combine funcs
func (fm *Manager) EnsureFileRO(path string) (*os.File, error) {
	cache, ok := fm.files[path]
	if ok {
		return cache, nil
	}

	// file not in cache
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	fm.files[path] = file
	return file, nil
}

func (b *Manager) close() {
	for _, v := range b.files {
		v.Close()
	}
}
