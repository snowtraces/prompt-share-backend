package storage

import "io"

type Storage interface {
	Save(filename string, data io.Reader) (path string, err error)
	Delete(path string) error
	Open(path string) (io.ReadCloser, error)
	Exists(path string) (bool, error)
}
