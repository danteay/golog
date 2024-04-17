package zerolog

import (
	"fmt"
	"io"
	"os"

	"github.com/danteay/golog/fields"
	"github.com/danteay/golog/levels"
	"github.com/rs/zerolog"
)

// Adapter is a zerolog adapter implementation
type Adapter struct {
	logger zerolog.Logger
}

func New(opts ...Option) *Adapter {
	logOpts := options{
		level:   levels.Info,
		colored: false,
	}

	for _, opt := range opts {
		opt(&logOpts)
	}

	if logOpts.logger != nil {
		return &Adapter{
			logger: *logOpts.logger,
		}
	}

	return &Adapter{
		logger: getLogger(logOpts),
	}
}

// Logger returns the zerolog logger instance
func (a *Adapter) Logger() zerolog.Logger {
	return a.logger
}

// Log logs a message with the given level, error, fields, and message
func (a *Adapter) Log(level levels.Level, err error, logFields *fields.Fields, msg string, args ...any) {
	log := a.getLog(level)

	if err != nil {
		log = log.Err(err)
	}

	if logFields != nil {
		for k, v := range logFields.Data() {
			log = log.Interface(k, v)
		}
	}

	log.Msg(fmt.Sprintf(msg, args...))
}

func (a *Adapter) getLog(level levels.Level) *zerolog.Event {
	events := map[levels.Level]func() *zerolog.Event{
		levels.Debug: a.logger.Debug,
		levels.Info:  a.logger.Info,
		levels.Warn:  a.logger.Warn,
		levels.Error: a.logger.Error,
		levels.Fatal: a.logger.Fatal,
		levels.Panic: a.logger.Panic,
	}

	event, exists := events[level]
	if !exists {
		return a.logger.Info()
	}

	return event()
}

func getLevels(level levels.Level) zerolog.Level {
	levelList := map[levels.Level]zerolog.Level{
		levels.Debug: zerolog.DebugLevel,
		levels.Info:  zerolog.InfoLevel,
		levels.Warn:  zerolog.WarnLevel,
		levels.Error: zerolog.ErrorLevel,
		levels.Fatal: zerolog.FatalLevel,
		levels.Panic: zerolog.PanicLevel,
	}

	zl, exists := levelList[level]
	if !exists {
		return zerolog.InfoLevel
	}

	return zl
}

func getLogger(opts options) zerolog.Logger {
	var writer io.Writer = os.Stdout

	if opts.colored {
		writer = zerolog.ConsoleWriter{Out: os.Stdout}
	}

	if opts.writer != nil {
		writer = opts.writer
	}

	return zerolog.New(writer).With().Timestamp().Logger().Level(getLevels(opts.level))
}
