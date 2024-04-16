package golog

import (
	"context"

	"github.com/danteay/golog/adapters/slog"
	"github.com/danteay/golog/fields"
	"github.com/danteay/golog/internal/contextfields"
	"github.com/danteay/golog/internal/errors"
	"github.com/danteay/golog/levels"
)

type Adapter interface {
	Log(level levels.Level, err error, logFields *fields.Fields, msg string, args ...any)
}

type Logger struct {
	ctx    context.Context
	logger Adapter
	fields *fields.Fields
	err    error
}

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

// Log logs a message with the provided level, message and arguments.
func (l *Logger) Log(level levels.Level, msg string, args ...any) {
	defer l.reset()

	l.fields.Merge(contextfields.Fields(l.ctx))

	if l.err != nil {
		l.fields.Set("stack", errors.GetStackTrace())
	}

	l.logger.Log(level, l.err, l.fields, msg, args...)
}

// Debug logs a message with the Debug level.
func (l *Logger) Debug(msg string, args ...any) {
	l.Log(levels.Debug, msg, args...)
}

// Info logs a message with the Info level.
func (l *Logger) Info(msg string, args ...any) {
	l.Log(levels.Info, msg, args...)
}

// Warn logs a message with the Warn level.
func (l *Logger) Warn(msg string, args ...any) {
	l.Log(levels.Warn, msg, args...)
}

// Error logs a message with the Error level.
func (l *Logger) Error(msg string, args ...any) {
	l.Log(levels.Error, msg, args...)
}

// Fatal logs a message with the Fatal level.
func (l *Logger) Fatal(msg string, args ...any) {
	l.Log(levels.Fatal, msg, args...)
}

// Panic logs a message with the Panic level.
func (l *Logger) Panic(msg string, args ...any) {
	l.Log(levels.Panic, msg, args...)
}

func (l *Logger) reset() {
	l.fields = fields.New()
	l.err = nil
}
