package para

import (
	"os"
)

type FileSystem struct {
	rootPath string
}

func NewFileSystem(s string) *FileSystem {
	return &FileSystem{s}
}

func (fs *FileSystem) Init() error {
	_, err := os.Stat(fs.rootPath)
	if err != nil {
		return err
	}
	return nil
}
