package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var logger *logrus.Logger

var Trace = func(args ...interface{}) {
	logger.Trace(args...)
}
var Debug = func(args ...interface{}) {
	logger.Debug(args...)
}
var Info = func(args ...interface{}) {
	logger.Info(args...)
}
var Warn = func(args ...interface{}) {
	logger.Warn(args...)
}
var Error = func(args ...interface{}) {
	logger.Error(args...)
}
var Fatal = func(args ...interface{}) {
	logger.Fatal(args...)
}
var Panic = func(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

var Tracef = func(format string, args ...interface{}) {
	logger.Tracef(format, args...)
}
var Debugf = func(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}
var Infof = func(format string, args ...interface{}) {
	logger.Infof(format, args...)
}
var Warnf = func(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}
var Errorf = func(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}
var Fatalf = func(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}
var Panicf = func(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

func init() {
	fmt.Println("log init")
	logger = logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			//fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
			//return frame.Function, fileName
			sps := strings.Split(frame.File, "golangrustnode")
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
		logger.SetLevel(logrus.InfoLevel) // 默认级别
	} else {
		logger.SetLevel(level)
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

func SetLevel(level string) {
	SetDebug()
	l, err := logrus.ParseLevel(level)
	if err != nil {
		Debug(err)
		return
	}
	logger.SetLevel(l)
}

func SetPanic() {
	logger.SetLevel(logrus.PanicLevel)
}
func SetFatal() {
	logger.SetLevel(logrus.FatalLevel)
}
func SetError() {
	logger.SetLevel(logrus.ErrorLevel)
}
func SetWarn() {
	logger.SetLevel(logrus.WarnLevel)
}
func SetInfo() {
	logger.SetLevel(logrus.InfoLevel)
}
func SetDebug() {
	logger.SetLevel(logrus.DebugLevel)
}
func SetTrace() {
	logger.SetLevel(logrus.TraceLevel)
}
