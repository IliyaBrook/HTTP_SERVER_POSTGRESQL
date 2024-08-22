package utils

import (
	"log"
	"path/filepath"
	"runtime"
)

func GetFilePath() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("Runtime get path error")
	}
	dir := filepath.Dir(filename)
	return dir
}
