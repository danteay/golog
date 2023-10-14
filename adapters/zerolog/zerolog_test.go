package zerolog

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/danteay/golog/fields"
	"github.com/danteay/golog/levels"
)

type testMsg struct {
	Level   string `json:"level"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Key1    string `json:"key1"`
	Key2    int    `json:"key2"`
}

func TestAdapter_Log(t *testing.T) {
	t.Run("should log message", func(t *testing.T) {
		var logOutput bytes.Buffer

		logger := New(WithLevel(levels.Debug), WithWriter(&logOutput))

		logFields := fields.New().SetMap(map[string]any{
			"key1": "value1",
			"key2": 42,
		})

		msg := "Test message"
		err := errors.New("test error")

		logger.Log(levels.Debug, err, logFields, msg)

		res := testMsg{}

		if errUnmarshall := json.Unmarshal(logOutput.Bytes(), &res); errUnmarshall != nil {
			t.Fatal(errUnmarshall)
		}

		assert.Equal(t, levels.Debug.String(), res.Level)
		assert.Equal(t, msg, res.Message)
		assert.Equal(t, logFields.Get("key1"), res.Key1)
		assert.Equal(t, logFields.Get("key2"), res.Key2)
		assert.Equal(t, err.Error(), res.Error)
	})

	t.Run("should log message with no fields", func(t *testing.T) {
		var logOutput bytes.Buffer

		logger := New(WithLevel(levels.Debug), WithWriter(&logOutput))

		msg := "Test message"
		err := errors.New("test error")

		logger.Log(levels.Debug, err, nil, msg)

		res := testMsg{}

		if errUnmarshall := json.Unmarshal(logOutput.Bytes(), &res); errUnmarshall != nil {
			t.Fatal(errUnmarshall)
		}

		assert.Equal(t, levels.Debug.String(), res.Level)
		assert.Equal(t, msg, res.Message)
		assert.Equal(t, err.Error(), res.Error)
	})

	t.Run("should log message with no error", func(t *testing.T) {
		var logOutput bytes.Buffer

		logger := New(WithLevel(levels.Debug), WithWriter(&logOutput))

		logFields := fields.New().SetMap(map[string]any{
			"key1": "value1",
			"key2": 42,
		})

		msg := "Test message"

		logger.Log(levels.Debug, nil, logFields, msg)

		res := testMsg{}

		if errUnmarshall := json.Unmarshal(logOutput.Bytes(), &res); errUnmarshall != nil {
			t.Fatal(errUnmarshall)
		}

		assert.Equal(t, levels.Debug.String(), res.Level)
		assert.Equal(t, msg, res.Message)
		assert.Equal(t, logFields.Get("key1"), res.Key1)
		assert.Equal(t, logFields.Get("key2"), res.Key2)
		assert.Equal(t, "", res.Error)
	})

	t.Run("should log message with no fields and no error", func(t *testing.T) {
		var logOutput bytes.Buffer

		logger := New(WithLevel(levels.Debug), WithWriter(&logOutput))

		msg := "Test message"

		logger.Log(levels.Debug, nil, nil, msg)

		res := testMsg{}

		if errUnmarshall := json.Unmarshal(logOutput.Bytes(), &res); errUnmarshall != nil {
			t.Fatal(errUnmarshall)
		}

		assert.Equal(t, levels.Debug.String(), res.Level)
		assert.Equal(t, msg, res.Message)
		assert.Equal(t, "", res.Error)
	})

	t.Run("should not log message under level", func(t *testing.T) {
		var logOutput bytes.Buffer

		logger := New(WithLevel(levels.Info), WithWriter(&logOutput))

		logFields := fields.New().SetMap(map[string]any{
			"key1": "value1",
			"key2": 42,
		})

		msg := "Test message"
		err := errors.New("test error")

		logger.Log(levels.Debug, err, logFields, msg)

		assert.Equal(t, "", logOutput.String())
	})
}

func TestGetLevels(t *testing.T) {
	tests := map[levels.Level]zerolog.Level{
		levels.Debug: zerolog.DebugLevel,
		levels.Info:  zerolog.InfoLevel,
		levels.Warn:  zerolog.WarnLevel,
		levels.Error: zerolog.ErrorLevel,
		levels.Fatal: zerolog.FatalLevel,
		levels.Panic: zerolog.PanicLevel,
	}

	for level, expected := range tests {
		actual := getLevels(level)
		if actual != expected {
			t.Errorf("Expected level to be %v, but got %v", expected, actual)
		}
	}

	// Test an unsupported level
	level := getLevels(levels.Level(42))
	if level != zerolog.InfoLevel {
		t.Errorf("Expected level to be Info (default), but got %v", level)
	}
}
