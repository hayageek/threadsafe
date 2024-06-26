package threadsafe

import (
	"sync"
)

// Array represents a thread-safe array.
// It uses a mutex to ensure that all operations are thread-safe.
type Array[T any] struct {
	data []T
	mu   sync.RWMutex
}

// NewArray creates a new thread-safe array with a given size.
// Example:
// 	arr := threadsafe.NewArray 
func NewArray[T any](size int) *Array[T] {
	return &Array[T]{data: make([]T, size)}
}

// Get retrieves the value at the given index.
// It returns the value and a boolean indicating whether the index was valid.
// Example:
// 	value, ok := arr.Get(2)
func (a *Array[T]) Get(index int) (T, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if index < 0 || index >= len(a.data) {
		var zero T
		return zero, false
	}
	return a.data[index], true
}

// Set sets the value at the given index.
// It returns a boolean indicating whether the operation was successful.
// Example:
// 	ok := arr.Set(2, 100)
func (a *Array[T]) Set(index int, value T) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if index < 0 || index >= len(a.data) {
		return false
	}
	a.data[index] = value
	return true
}

// Length returns the length of the array.
// Example:
// 	length := arr.Length()
func (a *Array[T]) Length() int {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return len(a.data)
}
