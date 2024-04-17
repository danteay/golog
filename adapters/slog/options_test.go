package slog

import (
	"bytes"
	"testing"

	"github.com/danteay/golog/levels"
)

func TestWithLevel(t *testing.T) {
	opts := &options{}
	WithLevel(levels.Debug)(opts)

	if opts.level != levels.Debug {
		t.Errorf("Expected level to be Debug, but got %v", opts.level)
	}
}

func TestWithWriter(t *testing.T) {
	opts := &options{}
	writer := &bytes.Buffer{}
	WithWriter(writer)(opts)

	if opts.writer != writer {
		t.Error("Expected writer to be the provided io.Writer, but it's not")
	}
}

func TestOptionChaining(t *testing.T) {
	opts := &options{}
	WithLevel(levels.Error)(opts)
	writer := &bytes.Buffer{}
	WithWriter(writer)(opts)

	if opts.level != levels.Error {
		t.Errorf("Expected level to be Error, but got %v", opts.level)
	}

	if opts.writer != writer {
		t.Error("Expected writer to be the provided io.Writer, but it's not")
	}
}
