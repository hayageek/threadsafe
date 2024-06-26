package threadsafe

import (
	"testing"
)

func TestSlice(t *testing.T) {
	// Create a new thread-safe slice
	slice := NewSlice[int]()

	// Test appending values
	for i := 0; i < 5; i++ {
		slice.Append(i * 10)
	}

	// Test getting values
	for i := 0; i < slice.Length(); i++ {
		value, ok := slice.Get(i)
		if !ok || value != i*10 {
			t.Errorf("Expected %d, got %d at index %d", i*10, value, i)
		}
	}

	// Test setting values
	for i := 0; i < slice.Length(); i++ {
		if !slice.Set(i, i*20) {
			t.Errorf("Failed to set value at index %d", i)
		}
	}

	// Test getting updated values
	for i := 0; i < slice.Length(); i++ {
		value, ok := slice.Get(i)
		if !ok || value != i*20 {
			t.Errorf("Expected %d, got %d at index %d", i*20, value, i)
		}
	}

	// Test getting value out of range
	_, ok := slice.Get(10)
	if ok {
		t.Error("Expected false for out-of-range index, got true")
	}

	// Test setting value out of range
	if slice.Set(10, 100) {
		t.Error("Expected false for out-of-range index, got true")
	}
}
