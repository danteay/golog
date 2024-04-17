package zerolog

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"strings"

	"github.com/rs/zerolog"

	"github.com/danteay/golog/fields"
	"github.com/danteay/golog/levels"
)

// Adapter is a zerolog adapter implementation
type Adapter struct {
	logger    zerolog.Logger
	level     levels.Level
	writer    io.Writer
	withTrace bool
}

func New(opts ...Option) *Adapter {
	logOpts := options{
		level:   levels.Info,
		colored: false,
		writer:  os.Stdout,
	}

	for _, opt := range opts {
		opt(&logOpts)
	}

	adapter := &Adapter{
		level:     logOpts.level,
		writer:    getWriter(logOpts.writer, logOpts.colored),
		withTrace: logOpts.withTrace,
	}

	adapter.logger = getLogger(adapter.level, adapter.writer)

	return adapter
}

// Writer returns the writer for the adapter
func (a *Adapter) Writer() io.Writer {
	return a.writer
}

// SetWriter sets the writer for the adapter
func (a *Adapter) SetWriter(w io.Writer) {
	a.writer = w
}

// Logger returns the zerolog logger instance
func (a *Adapter) Logger() zerolog.Logger {
	return a.logger
}

// Log logs a message with the given level, error, fields, and message
func (a *Adapter) Log(level levels.Level, err error, logFields *fields.Fields, msg string, args ...any) {
	log := a.getLog(level)

	addErrFields(level, err, log, a.withTrace)

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

func addErrFields(level levels.Level, err error, evt *zerolog.Event, withTrace bool) {
	if err == nil {
		return
	}

	evt.Err(err)

	if withTrace || level == levels.TraceLevel {
		evt.Interface("stack", getStackTrace())
	}
}

func getStackTrace() []string {
	stack := strings.ReplaceAll(string(debug.Stack()), "\t", "")
	return strings.Split(stack, "\n")
}

func getLevels(level levels.Level) zerolog.Level {
	levelList := map[levels.Level]zerolog.Level{
		levels.NoLevel:    zerolog.NoLevel,
		levels.Disabled:   zerolog.Disabled,
		levels.TraceLevel: zerolog.TraceLevel,
		levels.Debug:      zerolog.DebugLevel,
		levels.Info:       zerolog.InfoLevel,
		levels.Warn:       zerolog.WarnLevel,
		levels.Error:      zerolog.ErrorLevel,
		levels.Fatal:      zerolog.FatalLevel,
		levels.Panic:      zerolog.PanicLevel,
	}

	zl, exists := levelList[level]
	if !exists {
		return zerolog.InfoLevel
	}

	return zl
}

func getWriter(baseWriter io.Writer, colored bool) io.Writer {
	if colored {
		return zerolog.ConsoleWriter{Out: baseWriter}
	}

	return baseWriter
}

func getLogger(level levels.Level, writer io.Writer) zerolog.Logger {
	return zerolog.New(writer).
		With().Timestamp().
		Logger().
		Level(getLevels(level))
}
