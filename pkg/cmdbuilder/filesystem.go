package cmdbuilder

import (
	"io/fs"
	"os"
)

type FileSystem interface {
	ReadDir(name string) ([]fs.DirEntry, error)
	Stat(name string) (fs.FileInfo, error)
}

type FileSystemImpl int

func NewFileSystem() *FileSystemImpl {
	return new(FileSystemImpl)
}

func (*FileSystemImpl) ReadDir(name string) ([]fs.DirEntry, error) {
	return os.ReadDir(name)
}

func (*FileSystemImpl) Stat(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}
