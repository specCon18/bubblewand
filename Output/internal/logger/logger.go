package logger

import (
	"os"

	"github.com/charmbracelet/log"
)

var Log *log.Logger

func init() {
	Log = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    false,
		ReportTimestamp: true,
		Prefix:          "test-test",
		Level:           log.InfoLevel,
	})
}

