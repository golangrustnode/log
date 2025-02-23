package log

import (
	"github.com/sirupsen/logrus"
)

func AddHook(hook logrus.Hook) {
	logger.AddHook(hook)
}
