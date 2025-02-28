package log

import (
	log "github.com/sirupsen/logrus"
	"strings"
	"testing"
)

func TestSetLevel(t *testing.T) {
	t.Log(logger)
	SetLevel("info")
	Infof("info level: %v", "fuckyou")
	Debug("debug level")
	Debug("debug level")
	Error("err log")
	res := strings.Split("golangrustnode", "tt")
	t.Log(res)
	t.Log(len(res))
	log.Info("info log")
}
