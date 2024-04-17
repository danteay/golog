package fields

import (
	"sync"
)

type Fields struct {
	mutex *sync.Mutex
	data  map[string]any
}

// New creates a new Fields instance
func New() *Fields {
	return &Fields{
		mutex: &sync.Mutex{},
		data:  make(map[string]any),
	}
}

// Set sets a key-value pair in the fields
func (f *Fields) Set(key string, value any) *Fields {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	f.data[key] = value

	return f
}

// SetMap sets a map of key-value pairs in the fields
func (f *Fields) SetMap(fields map[string]any) *Fields {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	for k, v := range fields {
		f.data[k] = v
	}

	return f
}

// Get returns the value of a key in the fields
func (f *Fields) Get(key string) any {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return f.data[key]
}

// Delete removes a key from the fields
func (f *Fields) Delete(key string) *Fields {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	delete(f.data, key)

	return f
}

// Clear removes all keys from the fields
func (f *Fields) Clear() *Fields {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	f.data = make(map[string]any)

	return f
}

// Copy returns a copy of the fields
func (f *Fields) Copy() *Fields {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	data := make(map[string]any)
	for k, v := range f.data {
		data[k] = v
	}

	return &Fields{
		mutex: &sync.Mutex{},
		data:  data,
	}
}

// Merge merges the fields with another fields
func (f *Fields) Merge(fields *Fields) *Fields {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if fields == nil {
		return f
	}

	for k, v := range fields.data {
		f.data[k] = v
	}

	return f
}

// Len returns the number of keys in the fields
func (f *Fields) Len() int {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return len(f.data)
}

// IsEmpty returns true if the fields are empty
func (f *Fields) IsEmpty() bool {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return len(f.data) == 0
}

// Data returns the fields as a map
func (f *Fields) Data() map[string]any {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	data := make(map[string]any)
	for k, v := range f.data {
		data[k] = v
	}

	return data
}
