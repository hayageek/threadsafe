
# Threadsafe

A Go package providing thread-safe implementations of arrays, slices, maps, stack & queue using generics and type constraints.

## Installation

To install the package, use the following command:

```sh
go get github.com/hayageek/threadsafe
```

## Usage

This package provides thread-safe implementations for arrays, slices, maps, stack & queue. Below are usage examples and API lists for each of these data types.

### Thread-Safe Array

A thread-safe array with a fixed size.

#### APIs

- `NewArray(size int) *Array[T]` - Creates a new thread-safe array with the given size.
- `(*Array[T]) Get(index int) (T, bool)` - Retrieves the value at the given index.
- `(*Array[T]) Set(index int, value T) bool` - Sets the value at the given index.
- `(*Array[T]) Length() int` - Returns the length of the array.

#### Example

```go
package main

import (
    "fmt"
    "github.com/hayageek/threadsafe"
)

func main() {
    // Create a new thread-safe array with size 5
    arr := threadsafe.NewArray[int](5)

    // Set values in the array
    for i := 0; i < arr.Length(); i++ {
        arr.Set(i, i*10)
    }

    // Get values from the array
    for i := 0; i < arr.Length(); i++ {
        value, _ := arr.Get(i)
        fmt.Println(value)
    }
}
```

### Thread-Safe Slice

A dynamically-sized, thread-safe slice.

#### APIs

- `NewSlice() *Slice[T]` - Creates a new thread-safe slice.
- `(*Slice[T]) Append(value T)` - Appends a value to the slice.
- `(*Slice[T]) Get(index int) (T, bool)` - Retrieves the value at the given index.
- `(*Slice[T]) Set(index int, value T) bool` - Sets the value at the given index.
- `(*Slice[T]) Length() int` - Returns the length of the slice.

#### Example

```go
package main

import (
    "fmt"
    "github.com/hayageek/threadsafe"
)

func main() {
    // Create a new thread-safe slice
    slice := threadsafe.NewSlice[int]()

    // Append values to the slice
    for i := 0; i < 5; i++ {
        slice.Append(i * 10)
    }

    // Get values from the slice
    for i := 0; i < slice.Length(); i++ {
        value, _ := slice.Get(i)
        fmt.Println(value)
    }

    // Set values in the slice
    for i := 0; i < slice.Length(); i++ {
        slice.Set(i, i*20)
    }

    // Get updated values from the slice
    for i := 0; i < slice.Length(); i++ {
        value, _ := slice.Get(i)
        fmt.Println(value)
    }
}
```

### Thread-Safe Map

A thread-safe map for storing key-value pairs.

#### APIs

- `NewMap() *Map[K, V]` - Creates a new thread-safe map.
- `(*Map[K, V]) Get(key K) (V, bool)` - Retrieves the value associated with the key.
- `(*Map[K, V]) Set(key K, value V)` - Sets the value for the given key.
- `(*Map[K, V]) Delete(key K)` - Deletes the value associated with the key.
- `(*Map[K, V]) Length() int` - Returns the number of key-value pairs in the map.
- `(*Map[K, V]) Keys() []K` - Returns a slice of all keys present in the map.
- `(*Map[K, V]) Values() []V` - Returns a slice of all values present in the map.

#### Example

```go
package main

import (
    "fmt"
    "github.com/hayageek/threadsafe"
)

func main() {
    // Create a new thread-safe map
    m := threadsafe.NewMap[string, int]()

    // Set values in the map
    m.Set("one", 1)
    m.Set("two", 2)
    m.Set("three", 3)

    // Get values from the map
    value, _ := m.Get("one")
    fmt.Println(value)

    value, _ = m.Get("two")
    fmt.Println(value)

    // Delete a key from the map
    m.Delete("two")

    // Check if a key exists
    _, ok := m.Get("two")
    if !ok {
        fmt.Println("Key 'two' not found")
    }

    // Get the length of the map
    length := m.Length()
    fmt.Println("Length:", length)

    // Get all keys from the map
    keys := m.Keys()
    fmt.Println("Keys:", keys)

    // Get all values from the map
    values := m.Values()
    fmt.Println("Values:", values)
}
```


### Thread-Safe Stack

A thread-safe stack for safely adding and removing items.

### Methods

- `NewStack() *Stack`: Creates a new thread-safe stack.
- `Push(value interface{})`: Adds an element to the stack.
- `Pop() (interface{}, bool)`: Removes and returns an element from the stack. Returns `false` if the stack is empty.
- `Len() int`: Returns the number of elements in the stack.

### Example

```go
package main

import (
    "fmt"
    "github.com/hayageek/threadsafe"
)

func main() {
    stack := threadsafe.NewStack()

    stack.Push(10)
    stack.Push(20)
    stack.Push(30)

    fmt.Println("Stack length:", stack.Len())

    value, ok := stack.Pop()
    if ok {
        fmt.Println("Popped value:", value)
    } else {
        fmt.Println("Stack is empty")
    }

    fmt.Println("Stack length after pop:", stack.Len())
}
```


### Thread-Safe Queue

A thread-safe queue for safely adding and removing items.

#### Methods

- `NewQueue() *Queue`: Creates a new thread-safe queue.
- `Enqueue(value interface{})`: Adds an element to the queue.
- `Dequeue() (interface{}, bool)`: Removes and returns an element from the queue. Returns `false` if the queue is empty.
- `Len() int`: Returns the number of elements in the queue.

#### Queue Example

```go
package main

import (
    "fmt"
    "github.com/hayageek/threadsafe"
)

func main() {
    queue := threadsafe.NewQueue()

    queue.Enqueue(10)
    queue.Enqueue(20)
    queue.Enqueue(30)

    fmt.Println("Queue length:", queue.Len())

    value, ok := queue.Dequeue()
    if ok {
        fmt.Println("Dequeued value:", value)
    } else {
        fmt.Println("Queue is empty")
    }

    fmt.Println("Queue length after dequeue:", queue.Len())
}
```

//Added Comment
