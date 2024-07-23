package log

import (
	"fmt"
	"os"
	"testing"
)

func TestSetLevel(t *testing.T) {
	SetLevel("info")
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	SetPath(cwd)
	fmt.Println("Current working directory:", cwd)
	Info("info level")
	Debug("debug level")
}
