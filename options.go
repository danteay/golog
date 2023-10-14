package golog

type options struct {
	adapter Adapter
}

type Option func(*options)

// WithAdapter sets the adapter for the logger.
func WithAdapter(adapter Adapter) Option {
	return func(opts *options) {
		if adapter == nil {
			return
		}

		opts.adapter = adapter
	}
}
