package threadsafe

import (
	"testing"
)

func TestArray(t *testing.T) {
	// Create a new thread-safe array with size 5
	arr := NewArray[int](5)

	// Test setting values
	for i := 0; i < arr.Length(); i++ {
		if !arr.Set(i, i*10) {
			t.Errorf("Failed to set value at index %d", i)
		}
	}

	// Test getting values
	for i := 0; i < arr.Length(); i++ {
		value, ok := arr.Get(i)
		if !ok || value != i*10 {
			t.Errorf("Expected %d, got %d at index %d", i*10, value, i)
		}
	}

	// Test getting value out of range
	_, ok := arr.Get(10)
	if ok {
		t.Error("Expected false for out-of-range index, got true")
	}

	// Test setting value out of range
	if arr.Set(10, 100) {
		t.Error("Expected false for out-of-range index, got true")
	}
}
