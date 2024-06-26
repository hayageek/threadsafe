package threadsafe

import (
	"testing"
)

func TestMap(t *testing.T) {
	// Create a new thread-safe map
	m := NewMap[string, int]()

	// Test setting values
	m.Set("one", 1)
	m.Set("two", 2)
	m.Set("three", 3)

	// Test getting values
	value, ok := m.Get("one")
	if !ok || value != 1 {
		t.Errorf("Expected 1, got %d", value)
	}

	value, ok = m.Get("two")
	if !ok || value != 2 {
		t.Errorf("Expected 2, got %d", value)
	}

	// Test getting non-existent key
	_, ok = m.Get("four")
	if ok {
		t.Error("Expected false for non-existent key, got true")
	}

	// Test deleting a key
	m.Delete("two")
	_, ok = m.Get("two")
	if ok {
		t.Error("Expected false for deleted key, got true")
	}

	// Test length
	length := m.Length()
	if length != 2 {
		t.Errorf("Expected length 2, got %d", length)
	}

	// Test keys
	keys := m.Keys()
	expectedKeys := map[string]bool{"one": true, "three": true}
	for _, key := range keys {
		if !expectedKeys[key] {
			t.Errorf("Unexpected key %s", key)
		}
	}

	// Test values
	values := m.Values()
	expectedValues := map[int]bool{1: true, 3: true}
	for _, value := range values {
		if !expectedValues[value] {
			t.Errorf("Unexpected value %d", value)
		}
	}
}
