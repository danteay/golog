package zerolog

import (
	"io"

	"github.com/rs/zerolog"

	"github.com/danteay/golog/levels"
)

type options struct {
	level   levels.Level
	writer  io.Writer
	colored bool
	logger  *zerolog.Logger
}

// Option defines the signature for the options.
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

// Colored sets the logger to use colored output.
func Colored() Option {
	return func(opts *options) {
		opts.colored = true
	}
}

// WithLogger sets the logger for the adapter.
func WithLogger(logger zerolog.Logger) Option {
	return func(opts *options) {
		opts.logger = &logger
	}
}
