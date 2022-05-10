package log

import (
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
)

func init()  {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:"2006-01-02 15:03:04",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			//处理文件名
			fileName := path.Base(frame.File)
			return frame.Function, fileName
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