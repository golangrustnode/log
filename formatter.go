package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

type CustomFormatter struct {
	logrus.Formatter
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Adjusting the depth to get the caller's caller
	_, file, line, ok := runtime.Caller(9) // 2 levels up the stack (1 would give us the direct caller)
	logLevel := entry.Level.String()
	if !ok {
		file = "unknown"
		line = 0
	}
	fsplits := strings.Split(file, "github.com/")
	filename := fsplits[len(fsplits)-1]
	//Creating custom log format
	logMessage := fmt.Sprintf("	%s:%d %s %s  \n", filename, line, logLevel, entry.Message)
	return []byte(logMessage), nil
}
