package log

import (
	"github.com/sirupsen/logrus"
	"runtime"
	"strconv"
	"strings"
)

func init()  {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			//fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
			//return frame.Function, fileName
			sps := strings.Split(frame.File,"github.com")
			len := len(sps)
			return "", sps[len-1]+":"+strconv.Itoa(frame.Line)
		},
	})
}

var Trace=logrus.Trace
var Debug=logrus.Debug
var Info =logrus.Info
var Warn=logrus.Warn
var Error=logrus.Error
var Fatal=logrus.Fatal
var Panic=logrus.Panic

func SetLevel(level string)  {
	SetDebug()
	l,err := logrus.ParseLevel(level)
	if err != nil {
		Debug(err)
		return
	}
	logrus.SetLevel(l)
}

func SetPanic() {
	logrus.SetLevel(logrus.PanicLevel)
}
func SetFatal()  {
	logrus.SetLevel(logrus.FatalLevel)
}
func SetError()  {
	logrus.SetLevel(logrus.ErrorLevel)
}
func SetWarn()  {
	logrus.SetLevel(logrus.WarnLevel)
}
func SetInfo()  {
	logrus.SetLevel(logrus.InfoLevel)
}
func SetDebug()  {
	logrus.SetLevel(logrus.DebugLevel)
}
func SetTrace()  {
	logrus.SetLevel(logrus.TraceLevel)
}