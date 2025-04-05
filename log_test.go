package log

import (
	"strings"
	"testing"
)

func TestSetLevel(t *testing.T) {
	t.Log(logger)
	SetLevel("info")
	Infof("info level: %v", "fuckyou")
	Error("err log")
	res := strings.Split("golangrustnode", "tt")
	t.Log(res)
	t.Log(len(res))
	Info("info log")
	
}
