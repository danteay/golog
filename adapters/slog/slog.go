package slog

import (
	"fmt"
	"github.com/danteay/golog/fields"
	"github.com/danteay/golog/levels"
	"log/slog"
	"os"
)

type Adapter struct {
	logger *slog.Logger
}

func New(opts ...Option) *Adapter {
	logOpts := options{
		level:  levels.Info,
		writer: os.Stdout,
	}

	for _, opt := range opts {
		opt(&logOpts)
	}

	if logOpts.logger != nil {
		return &Adapter{
			logger: logOpts.logger,
		}
	}

	return &Adapter{
		logger: getSlogInstance(logOpts),
	}
}

func (a *Adapter) Logger() *slog.Logger {
	return a.logger
}

func (a *Adapter) Log(level levels.Level, err error, logFields *fields.Fields, msg string, args ...any) {
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

	if err != nil {
		lf = append(lf, slog.Any("error", err))
	}

	msg = fmt.Sprintf(msg, args...)

	switch level {
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

func getSlogInstance(opts options) *slog.Logger {
	if opts.logger != nil {
		return opts.logger
	}

	level := getLevels(opts.level)

	handler := slog.NewJSONHandler(opts.writer, &slog.HandlerOptions{
		AddSource: false,
		Level:     level,
	})

	return slog.New(handler)
}

func getLevels(level levels.Level) slog.Level {
	levelList := map[levels.Level]slog.Level{
		levels.Debug: slog.LevelDebug,
		levels.Info:  slog.LevelInfo,
		levels.Warn:  slog.LevelWarn,
		levels.Error: slog.LevelError,
		levels.Fatal: slog.LevelError,
		levels.Panic: slog.LevelError,
	}

	sl, exists := levelList[level]
	if !exists {
		return slog.LevelInfo
	}

	return sl
}
