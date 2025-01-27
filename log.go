package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			//fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
			//return frame.Function, fileName
			sps := strings.Split(frame.File, "github.com")
			len := len(sps)
			return "", sps[len-1] + ":" + strconv.Itoa(frame.Line)
		},
	})
	SetLogLevel()
}

func SetLogLevel() {
	logLevel := os.Getenv("PLOG_LEVEL")
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.SetLevel(logrus.InfoLevel) // 默认级别
	} else {
		logrus.SetLevel(level)
	}
}

// trace debug info warn error fatal  panic
func ChangeLogLevel(levelstr string) {
	level, err := logrus.ParseLevel(levelstr)
	if err != nil {
		logrus.SetLevel(logrus.InfoLevel) // 默认级别
	} else {
		logrus.SetLevel(level)
	}
}

var Trace = logrus.Trace
var Debug = logrus.Debug
var Info = logrus.Info
var Warn = logrus.Warn
var Error = logrus.Error
var Fatal = logrus.Fatal
var Panic = logrus.Panic

var Tracef = logrus.Tracef
var Debugf = logrus.Debugf
var Infof = logrus.Infof
var Warnf = logrus.Warnf
var Errorf = logrus.Errorf
var Fatalf = logrus.Fatalf
var Panicf = logrus.Panicf

func SetLevel(level string) {
	SetDebug()
	l, err := logrus.ParseLevel(level)
	if err != nil {
		Debug(err)
		return
	}
	logrus.SetLevel(l)
}

func SetPanic() {
	logrus.SetLevel(logrus.PanicLevel)
}
func SetFatal() {
	logrus.SetLevel(logrus.FatalLevel)
}
func SetError() {
	logrus.SetLevel(logrus.ErrorLevel)
}
func SetWarn() {
	logrus.SetLevel(logrus.WarnLevel)
}
func SetInfo() {
	logrus.SetLevel(logrus.InfoLevel)
}
func SetDebug() {
	logrus.SetLevel(logrus.DebugLevel)
}
func SetTrace() {
	logrus.SetLevel(logrus.TraceLevel)
}
