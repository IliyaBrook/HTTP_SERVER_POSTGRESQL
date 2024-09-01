package pkg

import (
	"os"
	"path/filepath"
)

func ResolvePath(paths ...string) string {
	var root, _ = os.Getwd()
	return filepath.Join(append([]string{root}, paths...)...)
}
