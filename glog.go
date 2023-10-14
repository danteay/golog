// Package golog is a simple logger that uses zerolog as the underlying logger.
//
// Basically it adds syntactic sugar to zerolog, to make it easier to use if it can be.
//
// Also, it adds some features that are not present in zerolog, like context fields. This context fields
// are fields that are added to every log message, and are stored globally. To store them, whe should add
// a value to the context that identifies the execution, e.g. a request id using the key stored on
// glog/context.ExecutionContextKey.
//
// This context isolation is useful when we want to add fields to every log message of a specific execution,
// and we have multiple executions running at the same time, e.g. multiple requests.
//
// If there is no context value stored, the context fields will not be added to the log messages.
//
// Example:
//
//	 	import (
//	 		"context"
//	 		"github.com/danteay/golog"
//			logcontext "github.com/danteay/golog/context"
//	 	)
//
//		func main() {
//			ctx := context.WithValue(context.Background(), logcontext.ExecutionContextKey, "some-exec-id")
//			logger := New(ctx, WithWriter(&logOutput), WithLevel(DebugLevel), Colored())
//
//			logger.SetContextFields(map[string]any{
//				"stage": "dev",
//				"app-name": "some-name",
//			})
//
//			logger.Info("Hello %s", "world")
//			// {"level":"info","message":"Hello world","stage":"dev","app-name":"some-name","time":"2020-01-01T00:00:00Z"}
//
//			logger.Warn("This is a warning")
//			// {"level":"warn","message":"This is a warning","stage":"dev","app-name":"some-name","time":"2020-01-01T00:00:00Z"}
//
//			logger.Error("This is an error")
//			// {"level":"error","message":"This is an error","stage":"dev","app-name":"some-name","time":"2020-01-01T00:00:00Z"}
//
//			logger.FlushContextFields()
//
//			logger.Info("Hello %s", "world")
//			// {"level":"info","message":"Hello world","time":"2020-01-01T00:00:00Z"}
//
//			logger.Warn("This is a warning")
//			// {"level":"warn","message":"This is a warning","time":"2020-01-01T00:00:00Z"}
//
//			logger.Error("This is an error")
//			// {"level":"error","message":"This is an error","time":"2020-01-01T00:00:00Z"}
//		}
package golog

import "github.com/danteay/golog/internal/contextfields"

// FlushAllContextFields removes all stored context fields.
func FlushAllContextFields() {
	contextfields.FlushAll()
}
