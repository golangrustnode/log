package log

import "testing"

func TestSetLogRorate(t *testing.T) {
	SetLogRorate("/Users/developer/workspace/go/src/github.com/golangrustnode/log")
	Info("TestSetLogRorate")
}
