package utils

import (
	"os"
	"path/filepath"
)

func Workdir(path ...string) string {
	binaryRootPath, _ := filepath.Abs(os.Args[0])
	dir := filepath.Dir(binaryRootPath)
	if len(path) > 0 {
		return filepath.Join(dir, path[0])
	}

	return dir
}
