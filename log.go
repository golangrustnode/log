package log

import (
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
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
	//fmt.Println("log init")
	logger = logrus.New()

	logrus.SetReportCaller(true)
	logger.SetReportCaller(true)
	mytextformatter := &logrus.TextFormatter{
		ForceColors:            true,
		DisableColors:          false,
		FullTimestamp:          true, // 显示完整时间戳
		TimestampFormat:        "2006-01-02 15:04:05",
		DisableLevelTruncation: true, // 禁止级别文本截断，如 WARN 不会被截断为 WRN
		PadLevelText:           true, // 日志级别文字对齐
		QuoteEmptyFields:       true,
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			// 仅返回文件名（或相对路径）和行号
			/*sps := strings.Split(frame.File, "golangrustnode/")
			len := len(sps)
			return "", sps[len-1] + ":" + strconv.Itoa(frame.Line)*
			*/
			pc := make([]uintptr, 10)
			n := runtime.Callers(11, pc) // 调整skip数可能需要微调
			frames := runtime.CallersFrames(pc[:n])

			var frame runtime.Frame
			more := true
			for i := 0; i < 2 && more; i++ {
				frame, more = frames.Next()
			}
			sps := strings.Split(frame.File, "golangrustnode/")
			len := len(sps)
			//filename := filepath.Base(frame.File)
			filename := sps[len-1]
			line := strconv.Itoa(frame.Line)
			funcName := filepath.Base(frame.Function)

			return funcName, filename + ":" + line

		},
	}
	/* myloggertextformatter := &logrus.TextFormatter{
		ForceColors:            true,
		DisableColors:          false,
		FullTimestamp:          true, // 显示完整时间戳
		TimestampFormat:        "2006-01-02 15:04:05",
		DisableLevelTruncation: true, // 禁止级别文本截断，如 WARN 不会被截断为 WRN
		PadLevelText:           true, // 日志级别文字对齐
		QuoteEmptyFields:       true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			sps := strings.Split(frame.File, "golangrustnode/")
			len := len(sps)
			return "", sps[len-1] + ":" + strconv.Itoa(frame.Line)
		},
	}
	*/

	/*myjsonformatter := logrus.JSONFormatter{
		PrettyPrint:     false, // prettify the JSON output
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "@message",
		},
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return trimFunc(frame.Function), fmt.Sprintf("%s:%d", trimPath(frame.File), frame.Line)
		},
	}*/
	logrus.SetFormatter(mytextformatter)
	logger.SetFormatter(mytextformatter)
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
