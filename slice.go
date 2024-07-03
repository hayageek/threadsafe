package threadsafe

import (
	"reflect"
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

// Values returns a copy of the slice's data as a regular slice.
// Example:
//
//	values := slice.Values()
func (s *Slice[T]) Values() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	dataCopy := make([]T, len(s.data))
	copy(dataCopy, s.data)
	return dataCopy
}

// Remove removes the element at the given index.
// It returns a boolean indicating whether the operation was successful.
// Example:
//
//	ok := slice.Remove(2)
func (s *Slice[T]) Remove(index int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if index < 0 || index >= len(s.data) {
		return false
	}
	s.data = append(s.data[:index], s.data[index+1:]...)
	return true
}

// Contains checks if the slice contains the specified value.
// Example:
//
//	contains := slice.Contains(10)
func (s *Slice[T]) Contains(value T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, v := range s.data {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

// Clear removes all elements from the slice.
// Example:
//
//	slice.Clear()
func (s *Slice[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = []T{}
}

// Insert inserts a value at the specified index.
// It returns a boolean indicating whether the operation was successful.
// Example:
//
//	ok := slice.Insert(2, 10)
func (s *Slice[T]) Insert(index int, value T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if index < 0 || index > len(s.data) {
		return false
	}
	s.data = append(s.data[:index], append([]T{value}, s.data[index:]...)...)
	return true
}

// Copy returns a new thread-safe slice that is a copy of the current slice.
// Example:
//
//	copySlice := slice.Copy()
func (s *Slice[T]) Copy() *Slice[T] {
	s.mu.RLock()
	defer s.mu.RUnlock()
	dataCopy := make([]T, len(s.data))
	copy(dataCopy, s.data)
	return &Slice[T]{data: dataCopy}
}
