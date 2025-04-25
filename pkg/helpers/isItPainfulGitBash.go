package helpers

import (
	"os"
	"runtime"
	"strings"
)

func IsBrokenGitBash() bool {
	return runtime.GOOS == "windows" && strings.Contains(os.Getenv("TERM"), "xterm")
}
