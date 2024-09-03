package utils

import (
	"path/filepath"
)

func AbsPathJoin(paths ...string) string {
	joiningPath := filepath.Join(paths...)
	absPath, _ := filepath.Abs(joiningPath)
	return absPath
}
