package main

import "fmt"

type Node[T comparable] struct {
	Value    T
	Next     *Node[T]
	Previous *Node[T]
}

func newNode[T comparable](value T) *Node[T] {
	return &Node[T]{
		Value:    value,
		Next:     nil,
		Previous: nil,
	}
}

type Queue[T comparable] struct {
	Head   *Node[T]
	Tail   *Node[T]
	length int
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{
		Head:   nil,
		Tail:   nil,
		length: 0,
	}
}

func (q *Queue[T]) Enqueue(value T) {
	newNode := newNode[T](value)

	q.length += 1
	if q.length == 0 {
		q.Head = newNode
		q.Tail = newNode
		return
	}

	q.Tail.Next = newNode
	q.Tail = newNode

	return
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var value T

	if q.Head == nil {
		return value, false
	}

	head := q.Head
	head.Next = nil

	q.Head = q.Head.Next

	q.length -= 1
	return head.Value, true
}

func (q *Queue[T]) Peek() T {
	var value T
	if q.length > 0 {
		return q.Head.Value
	}
	return value
}

func main() {
	queue := NewQueue[string]()

	queue.Enqueue("first")
	queue.Enqueue("second")
	queue.Enqueue("third")
	queue.Enqueue("forth")
	queue.Enqueue("fifth")
	queue.Enqueue("sixth")

	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()

	value := queue.Peek()
	fmt.Println("VALUE", value)
}
