package logger

import (
	"log/slog"
	"os"
)

// Init initializes the global structured logger instance.
func Init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}
