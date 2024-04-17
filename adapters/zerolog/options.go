package zerolog

import (
	"io"

	"github.com/danteay/golog/levels"
)

type options struct {
	level     levels.Level
	writer    io.Writer
	colored   bool
	withTrace bool
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

// WithTrace sets the error trace for the logger.
func WithTrace() Option {
	return func(opts *options) {
		opts.withTrace = true
	}
}
