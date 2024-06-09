package config

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
)

// Logger returns a configured instance of the Charmbracelet log.Logger.
// This logger is set to write to os.Stderr, with additional settings to report
// the caller's location and timestamp in the log messages.
// The time format is set to time.Kitchen, and the prefix is set to "NullTask 🚫".
func Logger() *log.Logger {
	// Editing the log settings to make it easily to debug and understanding
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "NullTask 🚫",
	})

	return logger
}
