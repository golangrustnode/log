package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

func SetLevelBYENV() {
	logLevel := os.Getenv("PLOG_LEVEL")
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return
	}
	logrus.SetLevel(level)
	if logger != nil {
		logger.SetLevel(level)
	}
}

func SetLevel(level string) {
	l, err := logrus.ParseLevel(level)
	if err != nil {
		return
	}
	logger.SetLevel(l)
	if logger != nil {
		logger.SetLevel(l)
	}
}
