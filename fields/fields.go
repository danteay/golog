package fields

import "sync"

type Fields struct {
	mutex *sync.Mutex
	data  map[string]any
}

func New() *Fields {
	return &Fields{
		mutex: &sync.Mutex{},
		data:  make(map[string]any),
	}
}

func (f *Fields) Set(key string, value any) *Fields {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	f.data[key] = value

	return f
}

func (f *Fields) SetMap(fields map[string]any) *Fields {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	for k, v := range fields {
		f.data[k] = v
	}

	return f
}

func (f *Fields) Get(key string) any {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return f.data[key]
}

func (f *Fields) Delete(key string) *Fields {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	delete(f.data, key)

	return f
}

func (f *Fields) Clear() *Fields {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	f.data = make(map[string]any)

	return f
}

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

func (f *Fields) Len() int {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return len(f.data)
}

func (f *Fields) IsEmpty() bool {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return len(f.data) == 0
}

func (f *Fields) Data() map[string]any {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	data := make(map[string]any)
	for k, v := range f.data {
		data[k] = v
	}

	return data
}
