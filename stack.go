package threadsafe

import (
	"sync"

	"github.com/golang-collections/collections/stack"
)

// Stack is a thread-safe stack.
type Stack struct {
	s  *stack.Stack
	mu sync.Mutex
}

// NewStack creates a new thread-safe stack.
func NewStack() *Stack {
	return &Stack{
		s: stack.New(),
	}
}

// Push adds an element to the stack.
func (s *Stack) Push(value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.s.Push(value)
}

// Pop removes and returns an element from the stack.
func (s *Stack) Pop() (interface{}, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.s.Len() == 0 {
		return nil, false
	}
	return s.s.Pop(), true
}

// Len returns the number of elements in the stack.
func (s *Stack) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.s.Len()
}

// Peek returns the element at the top of the stack without removing it.
// Example:
//
//	value, ok := s.Peek()
func (s *Stack) Peek() (interface{}, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.s.Len() == 0 {
		return nil, false
	}
	return s.s.Peek(), true
}

// IsEmpty checks if the stack is empty.
// Example:
//
//	isEmpty := s.IsEmpty()
func (s *Stack) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.s.Len() == 0
}

// Clear removes all elements from the stack.
// Example:
//
//	s.Clear()
func (s *Stack) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.s = stack.New()
}

// Values returns a slice of all elements in the stack.
// Example:
//
//	values := s.Values()
func (s *Stack) Values() []interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Create a temporary slice to hold the values
	values := make([]interface{}, 0, s.s.Len())

	// Temporarily pop all elements to capture them
	length := s.s.Len()
	for i := 0; i < length; i++ {
		value := s.s.Pop()
		values = append(values, value)
	}

	// Push the elements back to restore the original state
	for i := len(values) - 1; i >= 0; i-- {
		s.s.Push(values[i])
	}

	return values
}
