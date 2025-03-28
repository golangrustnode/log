package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
)

var logger *logrus.Logger

var Trace = func(args ...interface{}) {
	safelog(logrus.TraceLevel, args...)
}
var Debug = func(args ...interface{}) {
	safelog(logrus.DebugLevel, args...)
}
var Info = func(args ...interface{}) {
	safelog(logrus.InfoLevel, args...)
}
var Warn = func(args ...interface{}) {
	safelog(logrus.WarnLevel, args...)
}
var Error = func(args ...interface{}) {
	safelog(logrus.ErrorLevel, args...)
}
var Fatal = func(args ...interface{}) {
	safelog(logrus.FatalLevel, args...)
}
var Panic = func(args ...interface{}) {
	safelog(logrus.PanicLevel, args...)
}

var Tracef = func(format string, args ...interface{}) {
	safelogf(logrus.TraceLevel, format, args...)
}
var Debugf = func(format string, args ...interface{}) {
	safelogf(logrus.DebugLevel, format, args...)
}
var Infof = func(format string, args ...interface{}) {
	safelogf(logrus.InfoLevel, format, args...)
}
var Warnf = func(format string, args ...interface{}) {
	safelogf(logrus.WarnLevel, format, args...)
}
var Errorf = func(format string, args ...interface{}) {
	safelogf(logrus.ErrorLevel, format, args...)
}
var Fatalf = func(format string, args ...interface{}) {
	safelogf(logrus.FatalLevel, format, args...)
}
var Panicf = func(format string, args ...interface{}) {
	safelogf(logrus.PanicLevel, format, args...)
}

func init() {
	fmt.Println("log init")
	logger = logrus.New()

	logrus.SetReportCaller(true)
	logger.SetReportCaller(true)
	myjsonformatter := logrus.JSONFormatter{
		PrettyPrint:     true, // prettify the JSON output
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "@message",
		},
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return trimFunc(frame.Function), fmt.Sprintf("%s:%d", trimPath(frame.File), frame.Line)
		},
	}
	logrus.SetFormatter(&CustomFormatter{
		JSONFormatter: myjsonformatter,
	})
	logger.SetFormatter(
		&CustomFormatter{
			JSONFormatter: myjsonformatter,
		},
	)
	/*
		logrus.SetFormatter(&CustomFormatter{&logrus.JSONFormatter{
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				//fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
				//return frame.Function, fileName
				sps := strings.Split(frame.File, "golangrustnode")
				len := len(sps)
				return "", sps[len-1] + ":" + strconv.Itoa(frame.Line)
			}},
		})

		logger.SetFormatter(&CustomFormatter{&logrus.JSONFormatter{
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				//fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
				//return frame.Function, fileName
				sps := strings.Split(frame.File, "golangrustnode")
				len := len(sps)
				return "", sps[len-1] + ":" + strconv.Itoa(frame.Line)
			}},
		})
	*/
	SetLevelBYENV()
}
