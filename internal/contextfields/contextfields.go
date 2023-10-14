package contextfields

import (
	"context"
	"reflect"
	"sync"

	"github.com/danteay/golog/fields"
)

const (
	// ExecutionContextKey is the key used to find in context the execution id to store global fields.
	ExecutionContextKey = "logger:execution_context"

	// DefaultContextValue is the default value used for the execution id.
	DefaultContextValue = "default"
)

var (
	contextFields map[any]*fields.Fields
	mutex         *sync.Mutex
)

func init() {
	contextFields = make(map[any]*fields.Fields)
	mutex = &sync.Mutex{}
}

// SetFields sets global fields according the stored execution id on context.
func SetFields(ctx context.Context, newFields map[string]any) {
	ctxVal := getExecKey(ctx)

	mutex.Lock()
	defer mutex.Unlock()

	val, exists := contextFields[ctxVal]

	if !exists || val == nil {
		contextFields[ctxVal] = fields.New()
	}

	for key, value := range newFields {
		contextFields[ctxVal].Set(key, value)
	}
}

// Flush removes the global fields according the stored execution id on context.
func Flush(ctx context.Context) {
	ctxVal := getExecKey(ctx)

	mutex.Lock()
	defer mutex.Unlock()

	delete(contextFields, ctxVal)
}

// FlushAll removes all global fields.
func FlushAll() {
	mutex.Lock()
	defer mutex.Unlock()

	contextFields = make(map[any]*fields.Fields)
}

// Fields returns the global fields according the stored execution id on context.
func Fields(ctx context.Context) *fields.Fields {
	ctxVal := getExecKey(ctx)

	execs := []any{ctxVal}

	if !reflect.DeepEqual(ctxVal, DefaultContextValue) {
		execs = append(execs, DefaultContextValue)
	}

	mutex.Lock()
	defer mutex.Unlock()

	f := fields.New()

	for _, exec := range execs {
		if val, exists := contextFields[exec]; exists && val != nil {
			f.Merge(val)
		}
	}

	return f
}

func getExecKey(ctx context.Context) any {
	if ctx == nil {
		return nil
	}

	val := ctx.Value(ExecutionContextKey)

	if val == nil {
		return DefaultContextValue
	}

	return val
}
