package log

import (
	"testing"
)

func TestSetLevel(t *testing.T) {
	SetLevel("info")
	Infof("info level: %v", "fuckyou")
	Debug("debug level")
}
