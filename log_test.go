package log

import (
	"testing"
)

func TestSetLevel(t *testing.T) {
	SetLevel("info")
	//fmt.Println("Current working directory:", cwd)
	Info("info level")
	Info("hello world")
	//Debug("debug level")
}
