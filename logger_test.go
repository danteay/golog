package golog

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/danteay/golog/adapters/slog"
<<<<<<< HEAD
=======
	"github.com/danteay/golog/adapters/zerolog"
>>>>>>> e2cbf83 (fix: modify mod versions)
	"github.com/danteay/golog/internal/contextfields"
	"github.com/danteay/golog/levels"
)

type testMsg struct {
	Level   string   `json:"level"`
	Message string   `json:"msg"`
	Stack   []string `json:"stack"`
	Error   string   `json:"error"`
	Key1    string   `json:"key1"`
	Key2    int      `json:"key2"`
	CtxKey  string   `json:"ctx_key"`
}

func TestNewLogger(t *testing.T) {
	ctx := context.Background()

	adapter := slog.New()
	logger := New(WithAdapter(adapter)).SetContext(ctx)

	assert.IsType(t, &Logger{}, logger)
	assert.Equal(t, ctx, logger.ctx)
	assert.Equal(t, adapter, logger.logger)
}

func TestLoggerFields(t *testing.T) {
	logger := New()

	fields := map[string]any{
		"key1": "value1",
		"key2": 42,
	}

	logger.Fields(fields)

	for key, value := range fields {
		if fieldValue := logger.fields.Get(key); fieldValue != value {
			t.Errorf("Expected field %s to be %v, but got %v", key, value, fieldValue)
		}
	}
}

func TestLoggerErr(t *testing.T) {
	logger := New()

	err := errors.New("test error")

	logger.Err(err)

	if logger.err == nil {
		t.Error("Expected logger error to be non-nil, but it's nil")
	}
	if logger.err != err {
		t.Errorf("Expected logger error to be %v, but got %v", err, logger.err)
	}
}

func TestLoggerLog(t *testing.T) {
	t.Run("Test log message", func(t *testing.T) {
		var logOutput bytes.Buffer

		adapter := slog.New(slog.WithWriter(&logOutput), slog.WithLevel(levels.Debug))
		logger := New(WithAdapter(adapter)).SetContext(context.Background())

		level := levels.Debug
		msg := "Test message"
		err := errors.New("test error")
		ctxKey := ""
		fields := map[string]any{
			"key1": "value1",
			"key2": 42,
		}

		// Log the message with fields and an error
		logger.Fields(fields).Err(err).Log(level, msg)

		res := testMsg{}

		t.Log(logOutput.String())

		if errMarshal := json.Unmarshal(logOutput.Bytes(), &res); errMarshal != nil {
			t.Fatal(errMarshal)
		}

		assert.Equal(t, level.String(), strings.ToLower(res.Level))
		assert.Equal(t, msg, res.Message)
		assert.Equal(t, err.Error(), res.Error)
		assert.Equal(t, fields["key1"], res.Key1)
		assert.Equal(t, fields["key2"], res.Key2)
		assert.Equal(t, ctxKey, res.CtxKey)
		assert.NotEmpty(t, res.Stack)
	})

	t.Run("Test log message with context fields", func(t *testing.T) {
		var logOutput bytes.Buffer

		ctx := context.WithValue(context.Background(), contextfields.ExecutionContextKey, "some-exec-id")
		adapter := slog.New(slog.WithWriter(&logOutput), slog.WithLevel(levels.Debug))
		logger := New(WithAdapter(adapter)).SetContext(ctx)
		defer logger.FlushContextFields()

		level := levels.Debug
		msg := "Test message"
		err := errors.New("test error")
		ctxKey := "some context val"

		fields := map[string]any{
			"key1": "value1",
			"key2": 42,
		}

		contextFields := map[string]any{
			"ctx_key": ctxKey,
		}

		// Log the message with fields and an error
		logger.Fields(fields).SetContextFields(contextFields).Err(err).Log(level, msg)

		res := testMsg{}

		t.Log(logOutput.String())

		if errMarshal := json.Unmarshal(logOutput.Bytes(), &res); errMarshal != nil {
			t.Fatal(errMarshal)
		}

		assert.Equal(t, level.String(), strings.ToLower(res.Level))
		assert.Equal(t, msg, res.Message)
		assert.Equal(t, err.Error(), res.Error)
		assert.Equal(t, fields["key1"], res.Key1)
		assert.Equal(t, fields["key2"], res.Key2)
		assert.Equal(t, ctxKey, res.CtxKey)
		assert.NotEmpty(t, res.Stack)

		logOutput.Reset()

		logger.Log(level, msg)

		res = testMsg{}

		if errMarshal := json.Unmarshal(logOutput.Bytes(), &res); errMarshal != nil {
			t.Fatal(errMarshal)
		}

		assert.Equal(t, level.String(), strings.ToLower(res.Level))
		assert.Equal(t, msg, res.Message)
		assert.Equal(t, "", res.Error)
		assert.Equal(t, "", res.Key1)
		assert.Equal(t, 0, res.Key2)
		assert.Equal(t, ctxKey, res.CtxKey)
		assert.Empty(t, res.Stack)
	})
}
