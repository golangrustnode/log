package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func AddHook(hook logrus.Hook) {
	if logger != nil {
		logger.AddHook(hook)
		fmt.Println("add hook", hook)
		return
	}
	fmt.Println("logger is nil")
}
