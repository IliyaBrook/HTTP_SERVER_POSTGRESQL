package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func InitCustomLogger() {
	log.SetReportCaller(true)
	log.SetFormatter(&customFormatter{})
}

type customFormatter struct{}

func (f *customFormatter) Format(entry *log.Entry) ([]byte, error) {
	logMessage := entry.Message
	path := fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
	timestamp := entry.Time.Format("03:04:05")

	formatted := fmt.Sprintf(`{
    "time": "%s",
    "path": "%s",
    "message": "%s
}`, timestamp, path, logMessage)
	return []byte(formatted + "\n"), nil
}
