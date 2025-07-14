package logger

import (
	"context"
	"log/slog"
	"os"
)

var Logger *slog.Logger

// Init initializes the global logger with JSON output.
func Init() {
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
}

// WithContext returns a logger that includes values from the context (if any).
func WithContext(ctx context.Context) *slog.Logger {
	return Logger.With(slog.Group("context"))
}
