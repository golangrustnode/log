package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var logger *logrus.Logger

var Trace = logger.Trace
var Debug = logger.Debug
var Info = logger.Info
var Warn = logger.Warn
var Error = logger.Error
var Fatal = logger.Fatal
var Panic = logger.Panic

var Tracef = logger.Tracef
var Debugf = logger.Debugf
var Infof = logger.Infof
var Warnf = logger.Warnf
var Errorf = logger.Errorf
var Fatalf = logger.Fatalf
var Panicf = logger.Panicf

func init() {
	logger = logrus.New()
	Trace = logger.Trace
	Debug = logger.Debug
	Info = logger.Info
	Warn = logger.Warn
	Error = logger.Error
	Fatal = logger.Fatal
	Panic = logger.Panic

	Tracef = logger.Tracef
	Debugf = logger.Debugf
	Infof = logger.Infof
	Warnf = logger.Warnf
	Errorf = logger.Errorf
	Fatalf = logger.Fatalf
	Panicf = logger.Panicf
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
