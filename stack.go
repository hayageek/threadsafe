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
