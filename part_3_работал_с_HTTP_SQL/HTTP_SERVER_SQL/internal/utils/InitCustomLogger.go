package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"runtime"
)

func InitCustomLogger() {
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   false,
		TimestampFormat: "03:04:05 PM",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			callerPath := fmt.Sprintf("%s:%d", f.File, f.Line)
			return "", callerPath
		},
	})
	log.SetFormatter(&customFormatter{})
}

type customFormatter struct{}

func (f *customFormatter) Format(entry *log.Entry) ([]byte, error) {
	logMessage := entry.Message
	path := entry.Caller.File
	return []byte(fmt.Sprintf("%s %s\n", path, logMessage)), nil
}
