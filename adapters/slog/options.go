package slog

import (
	"io"
	"log/slog"

	"github.com/danteay/golog/levels"
)

type options struct {
	level   levels.Level
	writer  io.Writer
	logger  *slog.Logger
	handler slog.Handler
}

type Option func(*options)

// WithLevel sets the log level for the logger.
func WithLevel(level levels.Level) Option {
	return func(opts *options) {
		opts.level = level
	}
}

// WithWriter sets the writer for the logger.
func WithWriter(writer io.Writer) Option {
	return func(opts *options) {
		opts.writer = writer
	}
}

// WithLogger sets the logger for the adapter.
func WithLogger(logger *slog.Logger) Option {
	return func(opts *options) {
		opts.logger = logger
	}
}

// WithHandler sets the handler for the logger. If a handler is set, the writer and level options will be ignored.
func WithHandler(handler slog.Handler) Option {
	return func(opts *options) {
		opts.handler = handler
	}
}
