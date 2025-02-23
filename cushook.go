package log

import (
	"github.com/sirupsen/logrus"
)

func AddHook(hook logrus.Hook) {
	logrus.AddHook(hook)
}
