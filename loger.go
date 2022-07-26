package log

import (
	"context"
)

// A Logger represents a logger.
type Logger interface {
	// Error logs a message at error level.
	Error(format string, v ...interface{})
	// Infof logs a message at info level.
	Info(format string, v ...interface{})

	// Infof logs a message at info level.
	Debug(format string, v ...interface{})

	// WithContext returns a new logger with the given context.
	WithContext(context.Context) Logger
}
