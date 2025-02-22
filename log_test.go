package log

import (
	"strings"
	"testing"
)

func TestSetLevel(t *testing.T) {
	SetLevel("info")
	Infof("info level: %v", "fuckyou")
	Debug("debug level")
	res := strings.Split("golangrustnode", "tt")
	t.Log(res)
	t.Log(len(res))

}
