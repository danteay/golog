package levels

import (
	"strconv"
)

var (
	// TraceValue is the value used for the trace level field.
	TraceValue = "trace"
	// DebugValue is the value used for the debug level field.
	DebugValue = "debug"
	// InfoValue is the value used for the info level field.
	InfoValue = "info"
	// WarnValue is the value used for the warn level field.
	WarnValue = "warn"
	// ErrorValue is the value used for the error level field.
	ErrorValue = "error"
	// FatalValue is the value used for the fatal level field.
	FatalValue = "fatal"
	// PanicValue is the value used for the panic level field.
	PanicValue = "panic"
)

// Level defines log levels.
type Level int8

const (
	// NoLevel defines an absent log level.
	NoLevel Level = iota + 1
	// Disabled disables the logger.
	Disabled
	// TraceLevel defines trace log level.
	TraceLevel
	// Debug defines debug log level.
	Debug
	// Info defines info log level.
	Info
	// Warn defines warn log level.
	Warn
	// Error defines error log level.
	Error
	// Fatal defines fatal log level.
	Fatal
	// Panic defines panic log level.
	Panic
)

// String returns the string representation of the log level int
func (l Level) String() string {
	values := map[Level]string{
		NoLevel:    "",
		Disabled:   "disabled",
		TraceLevel: TraceValue,
		Debug:      DebugValue,
		Info:       InfoValue,
		Warn:       WarnValue,
		Error:      ErrorValue,
		Fatal:      FatalValue,
		Panic:      PanicValue,
	}

	if value, exists := values[l]; exists {
		return value
	}

	return strconv.Itoa(int(l))
}
