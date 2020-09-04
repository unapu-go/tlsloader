package tlsloader

import (
	"io"
	"os"
)

type Storage interface {
	Reader() (r io.ReadCloser, err error)
}

type PairStorage struct {
	Cert, Key Storage
}

type FileStorage struct {
	Path string
}

func (this *FileStorage) Reader() (io.ReadCloser, error) {
	return os.Open(this.Path)
}

func NewFileStorage(path string) *FileStorage {
	return &FileStorage{Path: path}
}

func NewSafeFilePairStorage(certPath, keyPath string, mode ...os.FileMode) *PairStorage {
	return &PairStorage{NewFileStorage(certPath), NewFileStorage(keyPath)}
}
