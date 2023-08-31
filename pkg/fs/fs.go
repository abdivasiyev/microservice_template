package fs

import (
	"path"
	"path/filepath"
	"runtime"

	"go.uber.org/fx"
)

var FxOption = fx.Provide(New)

type FS interface {
	RootDir() string
}

type fileSystem struct{}

func (f *fileSystem) RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(path.Dir(b)))
	return filepath.Dir(d)
}

func New() FS {
	return &fileSystem{}
}
