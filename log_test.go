package log

import (
	"testing"
)

func TestSetLevel(t *testing.T) {
	SetLevel("info")
	Info("info level")
	Debug("debug level")
}
