package main

import (
	"fmt"
	"testing"
)

// Node represents a node in the linked list
type Node[T any] struct {
	data T
	next *Node[T]
}

// LinkedList represents a linked list (FIFO queue)
type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
}

// Push adds a new node at the end of the linked list (enqueue)
func (list *LinkedList[T]) Push(data T) {
	newNode := &Node[T]{data: data}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		list.tail = newNode
	}
}

// Pop removes the node at the front of the linked list (dequeue) and returns its value
func (list *LinkedList[T]) Pop() (T, error) {
	var empty T
	if list.head == nil {
		return empty, fmt.Errorf("queue is empty")
	}
	data := list.head.data
	list.head = list.head.next
	if list.head == nil {
		list.tail = nil
	}
	return data, nil
}

// SliceQueue represents a FIFO queue using a slice
type SliceQueue[T any] struct {
	data []T
}

// Push adds a new element at the end of the slice (enqueue)
func (q *SliceQueue[T]) Push(data T) {
	q.data = append(q.data, data)
}

// Pop removes the element at the front of the slice (dequeue) and returns its value
func (q *SliceQueue[T]) Pop() (T, error) {
	var empty T
	if len(q.data) == 0 {
		return empty, fmt.Errorf("queue is empty")
	}
	data := q.data[0]
	q.data = q.data[1:]
	return data, nil
}

func BenchmarkLinkedListQueue(b *testing.B) {
	queue := LinkedList[int]{}
	for i := 0; i < b.N; i++ {
		queue.Push(i)
	}
	for i := 0; i < b.N; i++ {
		_, _ = queue.Pop()
	}
}

func BenchmarkSliceQueue(b *testing.B) {
	queue := SliceQueue[int]{}
	for i := 0; i < b.N; i++ {
		queue.Push(i)
	}
	for i := 0; i < b.N; i++ {
		_, _ = queue.Pop()
	}
}

func main() {
	// This is just a placeholder for the main function
	// Run `go test -bench=.` to execute the benchmarks
}
