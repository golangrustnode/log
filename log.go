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