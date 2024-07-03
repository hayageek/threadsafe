package threadsafe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSlice(t *testing.T) {
	slice := NewSlice[int]()
	assert.Equal(t, 0, slice.Length())
}

func TestSliceAppend(t *testing.T) {
	slice := NewSlice[int]()
	slice.Append(42)
	assert.Equal(t, 1, slice.Length())
	value, ok := slice.Get(0)
	assert.True(t, ok)
	assert.Equal(t, 42, value)
}

func TestSliceGetSet(t *testing.T) {
	slice := NewSlice[int]()
	slice.Append(1)
	slice.Append(2)
	slice.Append(3)
	ok := slice.Set(1, 42)
	assert.True(t, ok)
	value, ok := slice.Get(1)
	assert.True(t, ok)
	assert.Equal(t, 42, value)
}

func TestSliceGetInvalidIndex(t *testing.T) {
	slice := NewSlice[int]()
	value, ok := slice.Get(10)
	assert.False(t, ok)
	assert.Equal(t, 0, value)
}

func TestSliceSetInvalidIndex(t *testing.T) {
	slice := NewSlice[int]()
	ok := slice.Set(10, 42)
	assert.False(t, ok)
}

func TestSliceRemove(t *testing.T) {
	slice := NewSlice[int]()
	slice.Append(1)
	slice.Append(2)
	slice.Append(3)
	ok := slice.Remove(1)
	assert.True(t, ok)
	assert.Equal(t, 2, slice.Length())
	value, ok := slice.Get(1)
	assert.True(t, ok)
	assert.Equal(t, 3, value)
}

func TestSliceRemoveInvalidIndex(t *testing.T) {
	slice := NewSlice[int]()
	ok := slice.Remove(10)
	assert.False(t, ok)
}

func TestSliceContains(t *testing.T) {
	slice := NewSlice[int]()
	slice.Append(1)
	slice.Append(2)
	slice.Append(3)
	assert.True(t, slice.Contains(2))
	assert.False(t, slice.Contains(42))
}

func TestSliceClear(t *testing.T) {
	slice := NewSlice[int]()
	slice.Append(1)
	slice.Append(2)
	slice.Append(3)
	slice.Clear()
	assert.Equal(t, 0, slice.Length())
}

func TestSliceInsert(t *testing.T) {
	slice := NewSlice[int]()
	slice.Append(1)
	slice.Append(2)
	slice.Append(3)
	ok := slice.Insert(1, 42)
	assert.True(t, ok)
	assert.Equal(t, 4, slice.Length())
	value, ok := slice.Get(1)
	assert.True(t, ok)
	assert.Equal(t, 42, value)
}

func TestSliceInsertInvalidIndex(t *testing.T) {
	slice := NewSlice[int]()
	ok := slice.Insert(10, 42)
	assert.False(t, ok)
}

func TestSliceCopy(t *testing.T) {
	slice := NewSlice[int]()
	slice.Append(1)
	slice.Append(2)
	slice.Append(3)
	copySlice := slice.Copy()
	assert.Equal(t, slice.Length(), copySlice.Length())
	for i := 0; i < slice.Length(); i++ {
		origValue, _ := slice.Get(i)
		copyValue, _ := copySlice.Get(i)
		assert.Equal(t, origValue, copyValue)
	}
}

func TestSliceValues(t *testing.T) {
	slice := NewSlice[int]()
	slice.Append(1)
	slice.Append(2)
	slice.Append(3)
	values := slice.Values()
	assert.Equal(t, 3, len(values))
	assert.Equal(t, 1, values[0])
	assert.Equal(t, 2, values[1])
	assert.Equal(t, 3, values[2])
}
