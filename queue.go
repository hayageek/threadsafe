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
