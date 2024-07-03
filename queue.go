package threadsafe

import (
	"sync"

	"github.com/golang-collections/collections/queue"
)

// Queue is a thread-safe queue.
type Queue struct {
	q  *queue.Queue
	mu sync.Mutex
}

// NewQueue creates a new thread-safe queue.
func NewQueue() *Queue {
	return &Queue{
		q: queue.New(),
	}
}

// Enqueue adds an element to the queue.
func (q *Queue) Enqueue(value interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.q.Enqueue(value)
}

// Dequeue removes and returns an element from the queue.
func (q *Queue) Dequeue() (interface{}, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.q.Len() == 0 {
		return nil, false
	}
	return q.q.Dequeue(), true
}

// Len returns the number of elements in the queue.
func (q *Queue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.q.Len()
}

// Peek returns the element at the front of the queue without removing it.
// Example:
//
//	value, ok := q.Peek()
func (q *Queue) Peek() (interface{}, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.q.Len() == 0 {
		return nil, false
	}
	return q.q.Peek(), true
}

// IsEmpty checks if the queue is empty.
// Example:
//
//	isEmpty := q.IsEmpty()
func (q *Queue) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.q.Len() == 0
}

// Clear removes all elements from the queue.
// Example:
//
//	q.Clear()
func (q *Queue) Clear() {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.q = queue.New()
}

// Values returns a slice of all elements in the queue.
// Example:
//
//	values := q.Values()
func (q *Queue) Values() []interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()

	// Create a temporary slice to hold the values
	values := make([]interface{}, 0, q.q.Len())

	// Temporarily dequeue all elements to capture them
	length := q.q.Len()
	for i := 0; i < length; i++ {
		value := q.q.Dequeue()
		values = append(values, value)
		q.q.Enqueue(value) // Re-enqueue the element
	}

	return values
}
