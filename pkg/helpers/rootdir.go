package helpers

import (
	"path"
	"path/filepath"
	"runtime"
)

func RootDir() string {
	_, b, _, _ := runtime.Caller(2)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
