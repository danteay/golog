package golog

import (
	"context"
	"io"

	"github.com/danteay/golog/adapters/slog"
	"github.com/danteay/golog/fields"
	"github.com/danteay/golog/internal/contextfields"
	"github.com/danteay/golog/levels"
)

// Adapter is the interface that wraps the Log method.
type Adapter interface {
	Log(level levels.Level, err error, logFields *fields.Fields, msg string, args ...any)
	Writer() io.Writer
	SetWriter(w io.Writer)
	Level() levels.Level
	SetLevel(level levels.Level)
}

// Logger is the main struct that holds the logger instance.
type Logger struct {
	ctx    context.Context
	logger Adapter
	fields *fields.Fields
	err    error
}

var _ io.Writer = (*Logger)(nil)

// New creates a new Logger instance by using the provided context and options.
// If no options are provided, the default options will be used (level: Info, colored: false).
func New(opts ...Option) *Logger {
	logOpts := options{
		adapter: slog.New(),
	}

	for _, opt := range opts {
		opt(&logOpts)
	}

	return &Logger{
		ctx:    context.Background(),
		fields: fields.New(),
		logger: logOpts.adapter,
	}
}

// SetContext sets the context to be used in the logger instance to identify and group log fields by execution
func (l *Logger) SetContext(ctx context.Context) *Logger {
	if ctx == nil {
		return l
	}

	l.ctx = ctx
	return l
}

// Field adds a field to the logger instance.
func (l *Logger) Field(key string, value any) *Logger {
	l.fields.Set(key, value)
	return l
}

// Fields adds multiple fields to the logger instance.
func (l *Logger) Fields(fields map[string]any) *Logger {
	for k, v := range fields {
		l.fields.Set(k, v)
	}

	return l
}

// SetContextFields sets fields that should be printed in every log message.
func (l *Logger) SetContextFields(fields map[string]any) *Logger {
	contextfields.SetFields(l.ctx, fields)
	return l
}

// FlushContextFields removes all context fields from the logger instance.
func (l *Logger) FlushContextFields() *Logger {
	contextfields.Flush(l.ctx)
	return l
}

// Err sets the error to be logged.
func (l *Logger) Err(err error) *Logger {
	l.err = err
	return l
}

// Write user the writer configured o the adapter to write the logs.
func (l *Logger) Write(p []byte) (n int, err error) {
	return l.logger.Writer().Write(p)
}

// Writer returns the writer for the adapter instance.
func (l *Logger) Writer() io.Writer {
	return l.logger.Writer()
}

// SetWriter sets the writer for the adapter instance.
func (l *Logger) SetWriter(w io.Writer) {
	l.logger.SetWriter(w)
}

// Level returns the level for the adapter instance.
func (l *Logger) Level() levels.Level {
	return l.logger.Level()
}

// SetLevel sets the level for the adapter instance.
func (l *Logger) SetLevel(level levels.Level) {
	l.logger.SetLevel(level)
}
