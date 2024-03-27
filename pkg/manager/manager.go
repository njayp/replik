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

// TODO handle same file accessed for R and W
func (m *Manager) EnsureFile(path string, f func(name string) (*os.File, error)) (*os.File, error) {
	cache, ok := m.files[path]
	if ok {
		return cache, nil
	}

	// file not in cache
	file, err := f(path)
	if err != nil {
		return nil, err
	}

	m.files[path] = file
	return file, nil
}

func (m *Manager) Close() {
	for _, v := range m.files {
		v.Close()
	}
}
