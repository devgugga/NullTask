package config

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
)

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
