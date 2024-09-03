package utils

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func InitCustomLogger() {
	log.SetReportCaller(true)
	log.SetFormatter(&customFormatter{})
}

type customFormatter struct{}

func (f *customFormatter) Format(entry *log.Entry) ([]byte, error) {
	logMessage := entry.Message
	path := fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
	timestamp := entry.Time.Format("03:04:05 PM")

	var fields string
	if len(entry.Data) > 0 {
		fieldsData, _ := json.Marshal(entry.Data)
		fields = fmt.Sprintf("    %s\"fields\": %s%s", colorCyan, fieldsData, colorReset)
	}

	separator := ""
	if fields != "" {
		separator = ",\n" + fields
	}

	formatted := fmt.Sprintf(`{
    %s"time": "%s"%s,
    %s"path": "%s"%s%s
    %s"message": "%s"%s
}`,
		colorGreen, timestamp, colorReset,
		colorBlue, path, colorReset, separator,
		colorPurple, logMessage, colorReset,
	)

	if fields == "" {
		formatted = fmt.Sprintf(`{
	%s"time": "%s"%s,
	%s"path": "%s"%s
	%s"message": "%s"%s
}`,
			colorGreen, timestamp, colorReset,
			colorBlue, path, colorReset,
			colorPurple, logMessage, colorReset,
		)
	}

	return []byte(formatted + "\n"), nil
}
