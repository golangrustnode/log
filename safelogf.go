package log

import "github.com/sirupsen/logrus"

func safelogf(level logrus.Level, format string, args ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Error("log panic", err)
		}
	}()
	if logger == nil {
		switch level {
		case logrus.TraceLevel:
			logrus.Tracef(format, args...)
		case logrus.DebugLevel:
			logrus.Debugf(format, args...)
		case logrus.InfoLevel:
			logrus.Infof(format, args...)
		case logrus.WarnLevel:
			logrus.Warnf(format, args...)
		case logrus.ErrorLevel:
			logrus.Errorf(format, args...)
		case logrus.FatalLevel:
			logrus.Fatalf(format, args...)
		case logrus.PanicLevel:
			logrus.Panicf(format, args...)
		}
		return
	}
	switch level {
	case logrus.TraceLevel:
		logger.Tracef(format, args...)
	case logrus.DebugLevel:
		logger.Debugf(format, args...)
	case logrus.InfoLevel:
		logger.Infof(format, args...)
	case logrus.WarnLevel:
		logger.Warnf(format, args...)
	case logrus.ErrorLevel:
		logger.Errorf(format, args...)
	case logrus.FatalLevel:
		logger.Fatalf(format, args...)
	case logrus.PanicLevel:
		logger.Panicf(format, args...)
	}
}
