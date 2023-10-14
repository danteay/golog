package fields

import (
	"sync"
	"testing"
)

func TestSetAndGet(t *testing.T) {
	f := New()
	f.Set("key1", "value1")

	value := f.Get("key1")
	if value != "value1" {
		t.Errorf("Expected 'value1', but got %v", value)
	}
}

func TestDelete(t *testing.T) {
	f := New()
	f.Set("key1", "value1")
	f.Delete("key1")

	value := f.Get("key1")
	if value != nil {
		t.Errorf("Expected nil, but got %v", value)
	}
}

func TestClear(t *testing.T) {
	f := New()
	f.Set("key1", "value1")
	f.Clear()

	if f.Len() != 0 {
		t.Errorf("Expected 0, but got %v", f.Len())
	}
}

func TestCopy(t *testing.T) {
	f1 := New()
	f1.Set("key1", "value1")

	f2 := f1.Copy()
	f2.Set("key2", "value2")

	if f1.Len() != 1 {
		t.Errorf("Expected 1, but got %v", f1.Len())
	}

	if f2.Len() != 2 {
		t.Errorf("Expected 2, but got %v", f2.Len())
	}
}

func TestMerge(t *testing.T) {
	f1 := New()
	f1.Set("key1", "value1")

	f2 := New()
	f2.Set("key2", "value2")

	f1.Merge(f2)

	if f1.Len() != 2 {
		t.Errorf("Expected 2, but got %v", f1.Len())
	}
}

func TestLen(t *testing.T) {
	f := New()
	f.Set("key1", "value1")

	if f.Len() != 1 {
		t.Errorf("Expected 1, but got %v", f.Len())
	}
}

func TestIsEmpty(t *testing.T) {
	f := New()

	if !f.IsEmpty() {
		t.Error("Expected Fields to be empty, but it's not")
	}

	f.Set("key1", "value1")

	if f.IsEmpty() {
		t.Error("Expected Fields to not be empty, but it is")
	}
}

func TestData(t *testing.T) {
	f := New()
	f.Set("key1", "value1")
	f.Set("key2", "value2")

	data := f.Data()

	if len(data) != 2 {
		t.Errorf("Expected 2, but got %v", len(data))
	}
}

func TestConcurrentSetAndGet(t *testing.T) {
	f := New()
	const goroutines = 100
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			f.Set("key1", "value1")
			wg.Done()
		}()
	}

	wg.Wait()

	value := f.Get("key1")
	if value != "value1" {
		t.Errorf("Expected 'value1', but got %v", value)
	}
}

func TestConcurrentClear(t *testing.T) {
	f := New()
	const goroutines = 100
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			f.Clear()
			wg.Done()
		}()
	}

	wg.Wait()

	if f.Len() != 0 {
		t.Errorf("Expected 0, but got %v", f.Len())
	}
}

func TestConcurrentCopy(_ *testing.T) {
	f := New()
	const goroutines = 100
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			f.Copy()
			wg.Done()
		}()
	}

	wg.Wait()
}
