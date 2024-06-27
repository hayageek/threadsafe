package threadsafe

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()

	// Test Enqueue
	q.Enqueue(1)
	q.Enqueue(2)
	if q.Len() != 2 {
		t.Errorf("Expected length 2, got %d", q.Len())
	}

	// Test Dequeue
	val, ok := q.Dequeue()
	if !ok || val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}
	val, ok = q.Dequeue()
	if !ok || val != 2 {
		t.Errorf("Expected 2, got %v", val)
	}

	// Test Dequeue from empty queue
	val, ok = q.Dequeue()
	if ok || val != nil {
		t.Errorf("Expected nil, got %v", val)
	}
}
