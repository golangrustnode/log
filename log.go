package log

import (
	"github.com/sirupsen/logrus"
)

func init()  {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

var Trace=logrus.Trace
var Debug=logrus.Debug
var Info =logrus.Info
var Warn=logrus.Warn
var Error=logrus.Error
var Fatal=logrus.Fatal
var Panic=logrus.Panic

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