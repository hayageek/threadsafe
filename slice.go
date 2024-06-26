package threadsafe

import (
	"sync"
)

// Slice represents a thread-safe slice.
// It uses a mutex to ensure that all operations are thread-safe.
type Slice[T any] struct {
	data []T
	mu   sync.RWMutex
}

// NewSlice creates a new thread-safe slice.
// Example:
//
//	slice := threadsafe.NewSlice[int]()
func NewSlice[T any]() *Slice[T] {
	return &Slice[T]{data: []T{}}
}

// Append appends a value to the slice.
// Example:
//
//	slice.Append(10)
func (s *Slice[T]) Append(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data, value)
}

// Get retrieves the value at the given index.
// It returns the value and a boolean indicating whether the index was valid.
// Example:
//
//	value, ok := slice.Get(2)
func (s *Slice[T]) Get(index int) (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if index < 0 || index >= len(s.data) {
		var zero T
		return zero, false
	}
	return s.data[index], true
}

// Set sets the value at the given index.
// It returns a boolean indicating whether the operation was successful.
// Example:
//
//	ok := slice.Set(2, 100)
func (s *Slice[T]) Set(index int, value T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if index < 0 || index >= len(s.data) {
		return false
	}
	s.data[index] = value
	return true
}

// Length returns the length of the slice.
// Example:
//
//	length := slice.Length()
func (s *Slice[T]) Length() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}
