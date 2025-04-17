package log

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func SetLogRorate(workspace string) {
	if fileinfo, err := os.Stat(workspace); err != nil || !fileinfo.IsDir() {
		Errorf("dir %s is not legal: %v", workspace, err)
	}
	if !strings.HasSuffix(workspace, "/") {
		workspace = workspace + "/"
	}
	log := &lumberjack.Logger{
		Filename:   workspace + "pdaily.log", // log file path
		MaxSize:    10,                       // Maximum size in MB before rotating
		MaxBackups: 3,                        // Maximum number of old log files to retain
		MaxAge:     28,                       // Maximum number of days to retain old log files
		Compress:   true,                     // Compress old log files
	}
	logrus.SetOutput(log)
	if logger != nil {
		logger.SetOutput(log)
	}

}
