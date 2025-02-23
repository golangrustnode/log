package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func AddHook(hook logrus.Hook) {
	if logger != nil {
		logger.AddHook(hook)
		fmt.Println("add hook", hook)
	}
	fmt.Println("logger is nil")
}
