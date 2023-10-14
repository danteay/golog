package contextfields

import (
	"context"
	"sync"
	"testing"
)

func TestSetFields(t *testing.T) {
	t.Run("should set fields from execution context", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), ExecutionContextKey, "test_key")
		defer Flush(ctx)

		newFields := map[string]any{"key1": "value1", "key2": "value2"}

		SetFields(ctx, newFields)

		f := Fields(ctx)
		if f.Len() != 2 {
			t.Errorf("Expected 2 contextFields, but got %d", f.Len())
		}

		if f.Get("key1") != "value1" || f.Get("key2") != "value2" {
			t.Errorf("Fields not set correctly")
		}
	})

	t.Run("should set fields from default execution context", func(t *testing.T) {
		ctx := context.Background()
		defer Flush(ctx)

		newFields := map[string]any{"key1": "value1", "key2": "value2"}

		SetFields(ctx, newFields)

		f := Fields(ctx)
		if f.Len() != 2 {
			t.Errorf("Expected 2 contextFields, but got %d", f.Len())
		}

		if f.Get("key1") != "value1" || f.Get("key2") != "value2" {
			t.Errorf("Fields not set correctly")
		}
	})

	t.Run("should set fields to multiple execution contexts", func(t *testing.T) {
		ctx1 := context.WithValue(context.Background(), ExecutionContextKey, "test_key1")
		ctx2 := context.WithValue(context.Background(), ExecutionContextKey, "test_key2")
		ctx3 := context.Background()
		defer Flush(ctx1)
		defer Flush(ctx2)
		defer Flush(ctx3)

		SetFields(ctx1, map[string]any{"key1": "value1"})
		SetFields(ctx2, map[string]any{"key2": "value2"})
		SetFields(ctx3, map[string]any{"key3": "value3"})

		f1 := Fields(ctx1)
		if f1.Len() != 2 {
			t.Errorf("Expected 2 contextFields, but got %d", f1.Len())
		}

		if val := f1.Data()["key1"]; val != "value1" {
			t.Errorf("Expected value1, but got %v", val)
		}

		f2 := Fields(ctx2)
		if f2.Len() != 2 {
			t.Errorf("Expected 2 contextFields, but got %d", f2.Len())
		}

		if val := f2.Data()["key2"]; val != "value2" {
			t.Errorf("Expected value2, but got %v", val)
		}

		f3 := Fields(ctx3)
		if f3.Len() != 1 {
			t.Errorf("Expected 1 contextFields, but got %d", f3.Len())
		}

		if val := f3.Data()["key3"]; val != "value3" {
			t.Errorf("Expected value3, but got %v", val)
		}
	})
}

func TestFlush(t *testing.T) {
	ctx := context.WithValue(context.Background(), ExecutionContextKey, "test_key")

	newFields := map[string]any{"key1": "value1", "key2": "value2"}
	SetFields(ctx, newFields)

	Flush(ctx)

	f := Fields(ctx)
	if f.Len() > 0 {
		t.Error("Fields not flushed correctly")
	}
}

func TestFields(t *testing.T) {
	ctx := context.WithValue(context.Background(), ExecutionContextKey, "test_key")
	defer Flush(ctx)

	newFields := map[string]any{"key1": "value1", "key2": "value2"}
	SetFields(ctx, newFields)

	f := Fields(ctx)

	if f.Len() != 2 {
		t.Errorf("Expected 2 contextFields, but got %d", f.Len())
	}

	if f.Get("key1") != "value1" || f.Get("key2") != "value2" {
		t.Errorf("Fields not set correctly")
	}
}

func TestConcurrentSetFields(t *testing.T) {
	ctx := context.WithValue(context.Background(), ExecutionContextKey, "test_key")
	defer Flush(ctx)

	newFields := map[string]any{"key1": "value1"}

	const goroutines = 100
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			SetFields(ctx, newFields)
			wg.Done()
		}()
	}

	wg.Wait()

	f := Fields(ctx)
	t.Log(f)
	if f.Len() != 1 {
		t.Errorf("Expected 1 field, but got %d", f.Len())
	}
}

func TestConcurrentFlush(t *testing.T) {
	ctx := context.WithValue(context.Background(), ExecutionContextKey, "test_key")
	defer Flush(ctx)

	newFields := map[string]any{"key1": "value1"}
	SetFields(ctx, newFields)

	const goroutines = 100
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			Flush(ctx)
			wg.Done()
		}()
	}

	wg.Wait()

	f := Fields(ctx)
	if f.Len() > 0 {
		t.Error("Fields not flushed correctly in concurrent Flush")
	}
}
