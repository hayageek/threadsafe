package threadsafe

import (
	"sync"
)

// Map represents a thread-safe map.
// It uses a mutex to ensure that all operations are thread-safe.
type Map[K comparable, V any] struct {
	data map[K]V
	mu   sync.RWMutex
}

// NewMap creates a new thread-safe map.
// Example:
//
//	m := threadsafe.NewMap[string, int]()
func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{data: make(map[K]V)}
}

// Get retrieves the value associated with the key.
// It returns the value and a boolean indicating whether the key was found.
// Example:
//
//	value, ok := m.Get("key")
func (m *Map[K, V]) Get(key K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, exists := m.data[key]
	return value, exists
}

// Set sets the value for the given key.
// Example:
//
//	m.Set("key", 100)
func (m *Map[K, V]) Set(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

// GetOrSet returns the existing value for the key if present. Otherwise it sets
// defaultV for the key and returns it. The boolean return value indicates
// whether the key was already present.
// Example:
//
//	value, loaded := m.GetOrSet("key", 100)
func (m *Map[K, V]) GetOrSet(key K, defaultV V) (V, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if value, exists := m.data[key]; exists {
		return value, true
	}
	m.data[key] = defaultV
	return defaultV, false
}

// Delete removes the value associated with the key.
// Example:
//
//	m.Delete("key")
func (m *Map[K, V]) Delete(key K) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

// Pop removes the key from the map if it exists and returns the associated value.
// It returns the value and a boolean indicating whether the key was found.
// Example:
//
//	value, ok := m.Pop("key")
func (m *Map[K, V]) Pop(key K) (V, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	value, exists := m.data[key]
	if exists {
		delete(m.data, key)
	}
	return value, exists
}

// Length returns the number of key-value pairs in the map.
// Example:
//
//	length := m.Length()
func (m *Map[K, V]) Length() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.data)
}

// Keys returns a slice of all keys present in the map.
// Example:
//
//	keys := m.Keys()
func (m *Map[K, V]) Keys() []K {
	m.mu.RLock()
	defer m.mu.RUnlock()
	keys := make([]K, 0, len(m.data))
	for key := range m.data {
		keys = append(keys, key)
	}
	return keys
}

// Values returns a slice of all values present in the map.
// Example:
//
//	values := m.Values()
func (m *Map[K, V]) Values() []V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	values := make([]V, 0, len(m.data))
	for _, value := range m.data {
		values = append(values, value)
	}
	return values
}

// Contains checks if the map contains the specified key.
// Example:
//
//	contains := m.Contains("key")
func (m *Map[K, V]) Contains(key K) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	_, exists := m.data[key]
	return exists
}

// Clear removes all key-value pairs from the map.
// Example:
//
//	m.Clear()
func (m *Map[K, V]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data = make(map[K]V)
}

// Copy returns a new thread-safe map that is a copy of the current map.
// Example:
//
//	copyMap := m.Copy()
func (m *Map[K, V]) Copy() *Map[K, V] {
	m.mu.RLock()
	defer m.mu.RUnlock()
	dataCopy := make(map[K]V, len(m.data))
	for key, value := range m.data {
		dataCopy[key] = value
	}
	return &Map[K, V]{data: dataCopy}
}
