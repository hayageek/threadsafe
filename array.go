package threadsafe

import (
	"reflect"
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
//
//	arr := threadsafe.NewArray
func NewArray[T any](size int) *Array[T] {
	return &Array[T]{data: make([]T, size)}
}

// Get retrieves the value at the given index.
// It returns the value and a boolean indicating whether the index was valid.
// Example:
//
//	value, ok := arr.Get(2)
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
//
//	ok := arr.Set(2, 100)
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
//
//	length := arr.Length()
func (a *Array[T]) Length() int {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return len(a.data)
}

// Values returns a slice of all elements in the array.
// Example:
//
//	values := arr.Values()
func (a *Array[T]) Values() []T {
	a.mu.RLock()
	defer a.mu.RUnlock()
	dataCopy := make([]T, len(a.data))
	copy(dataCopy, a.data)
	return dataCopy
}

// Append appends a value to the array.
// Example:
//
//	arr.Append(10)
func (a *Array[T]) Append(value T) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.data = append(a.data, value)
}

// Remove removes the element at the given index.
// It returns a boolean indicating whether the operation was successful.
// Example:
//
//	ok := arr.Remove(2)
func (a *Array[T]) Remove(index int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if index < 0 || index >= len(a.data) {
		return false
	}
	a.data = append(a.data[:index], a.data[index+1:]...)
	return true
}

// Contains checks if the array contains the specified value.
// Example:
//
//	contains := arr.Contains(10)
func (a *Array[T]) Contains(value T) bool {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.data {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

// Clear removes all elements from the array.
// Example:
//
//	arr.Clear()
func (a *Array[T]) Clear() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.data = []T{}
}

// Insert inserts a value at the specified index.
// It returns a boolean indicating whether the operation was successful.
// Example:
//
//	ok := arr.Insert(2, 10)
func (a *Array[T]) Insert(index int, value T) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if index < 0 || index > len(a.data) {
		return false
	}
	a.data = append(a.data[:index], append([]T{value}, a.data[index:]...)...)
	return true
}

// Copy returns a new thread-safe array that is a copy of the current array.
// Example:
//
//	copyArray := arr.Copy()
func (a *Array[T]) Copy() *Array[T] {
	a.mu.RLock()
	defer a.mu.RUnlock()
	dataCopy := make([]T, len(a.data))
	copy(dataCopy, a.data)
	return &Array[T]{data: dataCopy}
}
