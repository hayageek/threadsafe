package threadsafe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	queue := NewQueue()
	assert.Equal(t, 0, queue.Len())
}

func TestQueueEnqueue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(42)
	assert.Equal(t, 1, queue.Len())
	value, ok := queue.Peek()
	assert.True(t, ok)
	assert.Equal(t, 42, value)
}

func TestQueueDequeue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(42)
	value, ok := queue.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 42, value)
	assert.Equal(t, 0, queue.Len())
}

func TestQueueDequeueEmpty(t *testing.T) {
	queue := NewQueue()
	value, ok := queue.Dequeue()
	assert.False(t, ok)
	assert.Nil(t, value)
}

func TestQueuePeek(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(42)
	value, ok := queue.Peek()
	assert.True(t, ok)
	assert.Equal(t, 42, value)
}

func TestQueuePeekEmpty(t *testing.T) {
	queue := NewQueue()
	value, ok := queue.Peek()
	assert.False(t, ok)
	assert.Nil(t, value)
}

func TestQueueLen(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(42)
	queue.Enqueue(43)
	assert.Equal(t, 2, queue.Len())
}

func TestQueueIsEmpty(t *testing.T) {
	queue := NewQueue()
	assert.True(t, queue.IsEmpty())
	queue.Enqueue(42)
	assert.False(t, queue.IsEmpty())
	queue.Dequeue()
	assert.True(t, queue.IsEmpty())
}

func TestQueueClear(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(42)
	queue.Enqueue(43)
	queue.Clear()
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
}

func TestQueueValues(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	values := queue.Values()
	assert.Equal(t, 3, len(values))
	assert.Equal(t, 1, values[0])
	assert.Equal(t, 2, values[1])
	assert.Equal(t, 3, values[2])
}
