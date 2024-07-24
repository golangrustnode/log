package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
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
			sps := strings.Split(frame.File, "github.com/golangrustnode/")
			len := len(sps)
			return "", sps[len-1] + ":" + strconv.Itoa(frame.Line)
		},
	})
}

var Trace = logrus.Trace
var Debug = logrus.Debug
var Info = logrus.Info
var Warn = logrus.Warn
var Error = logrus.Error
var Fatal = logrus.Fatal
var Panic = logrus.Panic

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

var stdFormatter *prefixed.TextFormatter
var fileFormatter *prefixed.TextFormatter

func SetPath(log_path string) {
	fileFormatter = &prefixed.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02.15:04:05.000000",
		ForceFormatting: true,
		ForceColors:     false,
		DisableColors:   true,
	}
	logName := fmt.Sprintf("%s/pcdnro.", log_path)
	writer, _ := rotatelogs.New(logName + "%Y%m%d")
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.DebugLevel: writer,
		logrus.ErrorLevel: writer,
	}, fileFormatter)
	logrus.SetOutput(os.Stdout)
	logrus.AddHook(lfHook)
}
