package levels

import "strconv"

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
	// Debug defines debug log level.
	Debug Level = iota
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
	// NoLevel defines an absent log level.
	NoLevel
	// Disabled disables the logger.
	Disabled

	// TraceLevel defines trace log level.
	TraceLevel Level = -1
	// Values less than TraceLevel are handled as numbers.
)

func (l Level) String() string {
	values := map[Level]string{
		TraceLevel: TraceValue,
		Debug:      DebugValue,
		Info:       InfoValue,
		Warn:       WarnValue,
		Error:      ErrorValue,
		Fatal:      FatalValue,
		Panic:      PanicValue,
		Disabled:   "disabled",
		NoLevel:    "",
	}

	if value, exists := values[l]; exists {
		return value
	}

	return strconv.Itoa(int(l))
}
