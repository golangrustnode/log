package log

import "github.com/sirupsen/logrus"

func safelog(level logrus.Level, args ...interface{}) {
	safelog2(level, args...)
}
func safelog2(level logrus.Level, args ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Error("log panic", err)
		}
	}()
	if logger == nil {
		switch level {
		case logrus.TraceLevel:
			logrus.Trace(args...)
		case logrus.DebugLevel:
			logrus.Debug(args...)
		case logrus.InfoLevel:
			logrus.Info(args...)
		case logrus.WarnLevel:
			logrus.Warn(args...)
		case logrus.ErrorLevel:
			logrus.Error(args...)
		case logrus.FatalLevel:
			logrus.Fatal(args...)
		case logrus.PanicLevel:
			logrus.Panic(args...)
		}
		return
	}
	switch level {
	case logrus.TraceLevel:
		logger.Trace(args...)
	case logrus.DebugLevel:
		logger.Debug(args...)
	case logrus.InfoLevel:
		logger.Info(args...)
	case logrus.WarnLevel:
		logger.Warn(args...)
	case logrus.ErrorLevel:
		logger.Error(args...)
	case logrus.FatalLevel:
		logger.Fatal(args...)
	case logrus.PanicLevel:
		logger.Panic(args...)
	}
}
