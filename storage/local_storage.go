package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	BasePath string
}

func NewLocalStorage(base string) *LocalStorage {
	os.MkdirAll(base, 0755)
	return &LocalStorage{BasePath: base}
}

func (s *LocalStorage) Save(filename string, data io.Reader) (string, error) {
	full := filepath.Join(s.BasePath, filename)
	dir := filepath.Dir(full)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	out, err := os.Create(full)
	if err != nil {
		return "", err
	}
	defer out.Close()
	if _, err := io.Copy(out, data); err != nil {
		return "", err
	}
	return full, nil
}

func (s *LocalStorage) Delete(path string) error {
	return os.Remove(path)
}

func (s *LocalStorage) Open(path string) (io.ReadCloser, error) {
	return os.Open(path)
}

func (s *LocalStorage) Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Helper to generate stored filename (simple)
func (s *LocalStorage) StoredName(prefix, filename string) string {
	return fmt.Sprintf("%s_%s", prefix, filepath.Base(filename))
}
