package threadsafe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMap(t *testing.T) {
	m := NewMap[string, int]()
	assert.Equal(t, 0, m.Length())
}

func TestMapSetGet(t *testing.T) {
	m := NewMap[string, int]()
	m.Set("key1", 42)
	value, ok := m.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, 42, value)
}

func TestMapGetNonExistentKey(t *testing.T) {
	m := NewMap[string, int]()
	value, ok := m.Get("nonexistent")
	assert.False(t, ok)
	assert.Equal(t, 0, value)
}

func TestMapDelete(t *testing.T) {
	m := NewMap[string, int]()
	m.Set("key1", 42)
	m.Delete("key1")
	value, ok := m.Get("key1")
	assert.False(t, ok)
	assert.Equal(t, 0, value)
}

func TestMapLength(t *testing.T) {
	m := NewMap[string, int]()
	m.Set("key1", 42)
	m.Set("key2", 43)
	assert.Equal(t, 2, m.Length())
}

func TestMapKeys(t *testing.T) {
	m := NewMap[string, int]()
	m.Set("key1", 42)
	m.Set("key2", 43)
	keys := m.Keys()
	assert.ElementsMatch(t, []string{"key1", "key2"}, keys)
}

func TestMapValues(t *testing.T) {
	m := NewMap[string, int]()
	m.Set("key1", 42)
	m.Set("key2", 43)
	values := m.Values()
	assert.ElementsMatch(t, []int{42, 43}, values)
}

func TestMapContains(t *testing.T) {
	m := NewMap[string, int]()
	m.Set("key1", 42)
	assert.True(t, m.Contains("key1"))
	assert.False(t, m.Contains("nonexistent"))
}

func TestMapClear(t *testing.T) {
	m := NewMap[string, int]()
	m.Set("key1", 42)
	m.Set("key2", 43)
	m.Clear()
	assert.Equal(t, 0, m.Length())
}

func TestMapCopy(t *testing.T) {
	m := NewMap[string, int]()
	m.Set("key1", 42)
	m.Set("key2", 43)
	copyMap := m.Copy()
	assert.Equal(t, m.Length(), copyMap.Length())
	for _, key := range m.Keys() {
		origValue, _ := m.Get(key)
		copyValue, _ := copyMap.Get(key)
		assert.Equal(t, origValue, copyValue)
	}
}
