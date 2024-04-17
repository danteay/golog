package slog

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime/debug"
	"strings"

	"github.com/danteay/golog/fields"
	"github.com/danteay/golog/levels"
)

// Adapter is a slog adapter implementation
type Adapter struct {
	logger    *slog.Logger
	level     levels.Level
	writer    io.Writer
	withTrace bool
}

func New(opts ...Option) *Adapter {
	logOpts := options{
		level:  levels.Info,
		writer: os.Stdout,
	}

	for _, opt := range opts {
		opt(&logOpts)
	}

	adapter := &Adapter{
		writer:    logOpts.writer,
		withTrace: logOpts.withTrace,
		level:     logOpts.level,
	}

	adapter.logger = getSlogInstance(adapter.level, adapter.writer)

	return adapter
}

func (a *Adapter) Writer() io.Writer {
	return a.writer
}

func (a *Adapter) SetWriter(w io.Writer) {
	a.writer = w
}

// Logger returns the slog logger instance
func (a *Adapter) Logger() *slog.Logger {
	return a.logger
}

// Log logs a message with the given level, error, fields, and message
func (a *Adapter) Log(level levels.Level, err error, logFields *fields.Fields, msg string, args ...any) {
	if level <= levels.Disabled {
		return
	}

	lenFields := 0
	if logFields != nil {
		lenFields = logFields.Len()
	}

	lf := make([]any, 0, lenFields)

	if logFields != nil {
		for key, val := range logFields.Data() {
			lf = append(lf, slog.Any(key, val))
		}
	}

	lf = getErrFields(level, err, lf, a.withTrace)

	msg = fmt.Sprintf(msg, args...)

	switch level {
	case levels.TraceLevel:
		a.logger.Debug(msg, lf...)
	case levels.Debug:
		a.logger.Debug(msg, lf...)
	case levels.Info:
		a.logger.Info(msg, lf...)
	case levels.Warn:
		a.logger.Warn(msg, lf...)
	case levels.Error:
		a.logger.Error(msg, lf...)
	case levels.Fatal:
		a.logger.Error(msg, lf...)
		os.Exit(1)
	case levels.Panic:
		a.logger.Error(msg, lf...)
		panic(err)
	default:
		a.logger.Info(msg, lf...)
	}
}

func getErrFields(level levels.Level, err error, curFields []any, withTrace bool) []any {
	if err == nil {
		return curFields
	}

	curFields = append(curFields, slog.Any("error", err))

	if level == levels.TraceLevel || withTrace {
		curFields = append(curFields, slog.Any("stack", getStackTrace()))
	}

	return curFields
}

func getStackTrace() []string {
	stack := strings.ReplaceAll(string(debug.Stack()), "\t", "")
	return strings.Split(stack, "\n")
}

func getSlogInstance(level levels.Level, writer io.Writer) *slog.Logger {
	handler := slog.NewJSONHandler(writer, &slog.HandlerOptions{
		AddSource: false,
		Level:     getLevels(level),
	})

	return slog.New(handler)
}

func getLevels(level levels.Level) slog.Level {
	levelList := map[levels.Level]slog.Level{
		levels.NoLevel:    slog.LevelInfo,
		levels.Disabled:   slog.LevelInfo,
		levels.TraceLevel: slog.LevelDebug,
		levels.Debug:      slog.LevelDebug,
		levels.Info:       slog.LevelInfo,
		levels.Warn:       slog.LevelWarn,
		levels.Error:      slog.LevelError,
		levels.Fatal:      slog.LevelError,
		levels.Panic:      slog.LevelError,
	}

	sl, exists := levelList[level]
	if !exists {
		return slog.LevelInfo
	}

	return sl
}
