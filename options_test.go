package golog

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/danteay/golog/adapters/slog"
)

func TestWithAdapter(t *testing.T) {
	opts := &options{}
	adapter := slog.New()

	WithAdapter(adapter)(opts)

	assert.Same(t, adapter, opts.adapter)
}
