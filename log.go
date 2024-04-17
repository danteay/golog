package golog

import (
	"github.com/danteay/golog/fields"
	"github.com/danteay/golog/internal/contextfields"
	"github.com/danteay/golog/internal/errors"
	"github.com/danteay/golog/levels"
)

// Log logs a message with the provided level, message and arguments.
func (l *Logger) Log(level levels.Level, msg string, args ...any) {
	defer l.reset()

	if level <= levels.Disabled {
		return
	}

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
