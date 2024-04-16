# golog

Syntactic sugar for logging and some other cool stuff.

Implements the Adapter pattern to replace the underlying logger. By default, it uses [slog][1], but you can create 
any other logger implementation by implementing the next interface:

```go
package golog

import (
	"github.com/danteay/golog/fields"
	"github.com/danteay/golog/levels"
)

type Adapter interface {
    Log(level levels.Level, err error, logFields *fields.Fields, msg string, args ...any)
}
```

## Requirements

- Go 1.22+

## Install

```bash
go get github.com/danteay/golog
```

## Basic usage

### Default Logger

```go
package main

import (
	"github.com/danteay/golog"
)

func main() {
	logger := golog.New() // Set a default logger using zerolog
	logger.Info("Hello world!")
}
```

### SetLog Fields

```go
package main

import (
	"errors"
	"github.com/danteay/golog"
)

func main() {
	logger := golog.New() // Set a default logger using zerolog

	// Single filed
	logger.Field("key", "value").Info("Hello world!")

	// Multi field
	logger.Fields(map[string]any{"key": "value", "key2": "value2"}).Info("Hello world!")

	// Set error
	logger.Err(errors.New("error")).Info("Hello world!")
}
```

## Configuring Zerolog adapter

### Current built-in options

```go
package main

import (
	"os"
	
    "github.com/danteay/golog"
    "github.com/danteay/golog/levels"
    "github.com/danteay/golog/adapters/zerolog"
)

func main() {
	adapter := zerolog.New(
		zerolog.WithLevel(levels.Debug),
		zerolog.Colored(),
	)
	
	logger := golog.New(golog.WithAdapter(adapter))
	
	// regular logging
}
```

| Option               | Description                                                                                                 | Default                  |
|----------------------|-------------------------------------------------------------------------------------------------------------|--------------------------|
| `zerolog.WithLevel`  | Sets the minimum logging level to the `zerolog` instance.                                                   | `levels.Debug`           |
| `zerolog.Colored`    | Set configuration to use a colored logging format. This is useful for local environments.                 | `false`                  |
| `zerolog.WithWriter` | Set a specific writer apart from the standard and colored outputs. If this option is used at the same time as the `Colored` option, it will override to use this new specific writer. | `null` |
| `zerolog.WithLogger` | Sets a preconfigured `zerolog.Logger` instance to use it on the adapter. If this option is set, it will omit any other option used to configure the adapter. | `null` |

## Working with context fields

Context fields is a concept added on this package to store log fields that should be added to every log entry. This is
useful when you want to add some fields to every log entry, but you don't want to add them manually every time.

```go
package main

import (
    "github.com/danteay/golog"
)

func main() {
	logger := golog.New()
	logger.SetContextFields(map[string]any{"key": "value"})
	
	logger.Info("Hello world!")
	// Output: {"level":"info","time":"2021-08-22T20:00:00-05:00","message":"Hello world!","key":"value"}
	
	logger.Warn("Hello world!")
	// Output: {"level":"warn","time":"2021-08-22T20:00:00-05:00","message":"Hello world!","key":"value"}
}
```

If you want to remove the context fields, you can use the `FlushContextFields` method.

```go
package main

import (
    "github.com/danteay/golog"
)

func main() {
	logger := golog.New()
	logger.SetContextFields(map[string]any{"key": "value"})
	
	logger.Info("Hello world!")
	// Output: {"level":"info","time":"2021-08-22T20:00:00-05:00","message":"Hello world!","key":"value"}
	
	logger.FlushContextFields()
	
	logger.Warn("Hello world!")
	// Output: {"level":"warn","time":"2021-08-22T20:00:00-05:00","message":"Hello world!"}
}
```

This implementation stores the context field in a default store, so any logger created on this way will share the same
context fields. 

If you want to create a logger with a different context fields, you should pass configure a `context.Context` instance
with the key `golog.ExecutionContextKey` and a unique value that refers to the execution of the logger, for exemple
a request ID.

```go
package main

import (
	"context"
    "github.com/danteay/golog"
)

func main() {
	logger := golog.New()
	logger.SetContextFields(map[string]any{"key": "value"})
	
	logger.Info("Hello world!")
	// Output: {"level":"info","time":"2021-08-22T20:00:00-05:00","message":"Hello world!","key":"value"}
	
	ctx := context.WithValue(context.Background(), golog.ExecutionContextKey, "request-id")
	custom := golog.New().SetContext(ctx)
	custom.SetContextFields(map[string]any{"key": "custom"})
	
	custom.Warn("Hello world!")
	// Output: {"level":"warn","time":"2021-08-22T20:00:00-05:00","message":"Hello world!","key":"custom"}
}
```

This implementation allows you to flush the context fields for a specific logger instance, without affecting the other.

```go
package main

import (
	"context"
	"github.com/danteay/golog"
)

func main() {
	logger := golog.New()
	logger.SetContextFields(map[string]any{"key": "value"})

	logger.Info("Hello world!")
	// Output: {"level":"info","time":"2021-08-22T20:00:00-05:00","message":"Hello world!","key":"value"}

	ctx := context.WithValue(context.Background(), golog.ExecutionContextKey, "request-id")
	custom := golog.New().SetContext(ctx)
	custom.SetContextFields(map[string]any{"key": "custom"})

	custom.Warn("Hello world!")
	// Output: {"level":"warn","time":"2021-08-22T20:00:00-05:00","message":"Hello world!","key":"custom"}

	custom.FlushContextFields()

	custom.Warn("Hello world!")
	// Output: {"level":"warn","time":"2021-08-22T20:00:00-05:00","message":"Hello world!"}
	logger.Warn("Hello world!")
	// Output: {"level":"warn","time":"2021-08-22T20:00:00-05:00","message":"Hello world!","key":"value"}
}
```

And if you want to flush all stored context fields no matter the logger instance, you can use the 
`golog.FlushAllContextFields` method.

## Caveats

### Memory leaks

Using context fields with no control may result on memory leaks. This is because the context fields are stored in a
global variable that is not flushed automatically, you should not use context fields if you don't really need them.

If you use context fields with the default store you should not add new fields regularly, you should add them once at
the beginning of the execution and flush them at the end of the execution.

If you use context fields with a custom store, you should flush them manually when you don't need them anymore, or
when the current execution scope ends.

[1]: https://betterstack.com/community/guides/logging/logging-in-go/
