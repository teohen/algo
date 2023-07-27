package main

import (
	"fmt"
)

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

type LinkedList[T comparable] struct {
	Head   *Node[T]
	Tail   *Node[T]
	length int
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{
		Head:   nil,
		length: 0}
}

func (l *LinkedList[T]) append(value T) {
	newNode := newNode[T](value)

	if l.length == 0 {
		l.Head = newNode
	}

	if l.Tail != nil {
		newNode.Previous = l.Tail
		l.Tail.Next = newNode
	}

	l.Tail = newNode

	l.length += 1
}

func (l *LinkedList[T]) getByIndex(idx int) (T, bool) {
	var t T

	if idx < 0 || idx > l.length-1 {
		return t, false
	}

	curr := l.Head

	if curr == nil {
		return t, false
	}

	idxCount := 0
	for curr != nil {

		if idxCount == idx {
			return curr.Value, true
		}

		idxCount += 1
		curr = curr.Next
	}
	return t, false
}

func (l *LinkedList[T]) prepend(value T) {
	newNode := newNode[T](value)

	if l.Head != nil {
		newNode.Next = l.Head
	}

	l.Head = newNode

	if l.Tail == nil {
		l.Tail = newNode
	}

	l.length += 1
}

func (l *LinkedList[T]) insertAt(value T, idx int) {
	newNode := newNode[T](value)

	if idx < 0 || idx > l.length-1 {
		return
	}

	if idx == 0 {
		l.prepend(value)
		return
	}

	if idx == l.length-1 {
		l.append(value)
		return
	}

	curr := l.Head
	count := -1
	for curr != nil {
		count += 1

		if count == idx {
			newNode.Next = curr.Next
			newNode.Previous = curr
			curr.Next.Previous = newNode
			curr.Next = newNode
			l.length += 1
			return
		}
		curr = curr.Next
	}
	return
}

func (l *LinkedList[T]) printAll() {
	curr := l.Head
	fmt.Println("LIST LENGTH", l.length)
	idx := -1
	for curr != nil {
		idx += 1
		fmt.Println("CUR", curr)
		fmt.Println("CUR.VALUE", curr.Value)
		fmt.Println("CUR.IDX", idx)

		curr = curr.Next
	}

	fmt.Println("HEAD", l.Head)
	fmt.Println("TAIL", l.Tail)
}
func (l *LinkedList[T]) getLength() int {
	return l.length
}

func (l *LinkedList[T]) Remove(value T) (T, bool) {
	curr := l.Head

	if value == l.Head.Value {
		l.Head = curr.Next
		l.length -= 1
		return curr.Value, true
	}

	for curr != nil {
		if curr.Value == value {
			if curr == l.Tail {
				l.Tail.Previous.Next = nil
				l.Tail = curr.Previous
				l.length -= 1
				return curr.Value, true
			}
			curr.Previous.Next = curr.Next
			curr.Next.Previous = curr.Previous

			curr.Next = nil
			curr.Previous = nil
			l.length -= 1
			return curr.Value, true
		}
		curr = curr.Next
	}
	return value, false
}

func (l *LinkedList[T]) RemoveAt(idx int) (T, bool) {
	var t T

	curr := l.Head

	if idx == 0 {
		t = l.Head.Value
		l.Head.Next.Previous = nil
		l.Head = l.Head.Next
		l.length -= 1
		return t, true
	}

	count := -1
	for curr != nil {
		count += 1

		if count == idx {
			t = curr.Value
			curr.Previous.Next = curr.Next
			curr.Next.Previous = curr.Previous
			l.length -= 1
			return t, true
		}

		curr = curr.Next
	}

	return t, false

}

func main() {
	list := NewLinkedList[string]()

	list.append("first")
	list.append("third")
	list.append("forth")
	list.insertAt("second", 1)
	list.printAll()

}
