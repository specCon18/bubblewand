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
		Prefix:          "bubblewand",
		Level:           log.InfoLevel,
	})
}
func initLogging() {
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		Log.Warn("Invalid log level; defaulting to info", "input", logLevel)
		level = log.InfoLevel
	}
	Log.SetLevel(level)
}
