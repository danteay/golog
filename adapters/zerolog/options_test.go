package zerolog

import (
	"bytes"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

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

func TestColored(t *testing.T) {
	opts := &options{}
	Colored()(opts)

	if !opts.colored {
		t.Error("Expected colored to be true, but it's not")
	}
}

func TestWithLogger(t *testing.T) {
	opts := &options{}
	WithLogger(zerolog.New(os.Stdout))(opts)

	assert.NotNil(t, opts.logger)
}

func TestOptionChaining(t *testing.T) {
	opts := &options{}
	WithLevel(levels.Error)(opts)
	writer := &bytes.Buffer{}
	WithWriter(writer)(opts)
	Colored()(opts)

	if opts.level != levels.Error {
		t.Errorf("Expected level to be Error, but got %v", opts.level)
	}

	if opts.writer != writer {
		t.Error("Expected writer to be the provided io.Writer, but it's not")
	}

	if !opts.colored {
		t.Error("Expected colored to be true, but it's not")
	}
}
