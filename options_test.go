package golog

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/danteay/golog/adapters/zerolog"
)

func TestWithAdapter(t *testing.T) {
	opts := &options{}
	adapter := zerolog.New()

	WithAdapter(adapter)(opts)

	assert.Same(t, adapter, opts.adapter)
}
