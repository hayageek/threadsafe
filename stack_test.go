package threadsafe

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()

	// Test Push
	s.Push(1)
	s.Push(2)
	if s.Len() != 2 {
		t.Errorf("Expected length 2, got %d", s.Len())
	}

	// Test Pop
	val, ok := s.Pop()
	if !ok || val != 2 {
		t.Errorf("Expected 2, got %v", val)
	}
	val, ok = s.Pop()
	if !ok || val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}

	// Test Pop from empty stack
	val, ok = s.Pop()
	if ok || val != nil {
		t.Errorf("Expected nil, got %v", val)
	}
}
