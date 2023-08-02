package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testError(t *testing.T, field string, expected any, got any) {
	t.Errorf("Field %s expected to be '%v' but got '%s'", field, expected, got)
}

func traverseList(list LinkedList[string], idx int) *Node[string] {
	curr := list.Head

	count := -1
	for curr != nil {
		count += 1

		if count == idx {
			return curr
		}
		return curr
	}
	return nil
}

func TestLinkedListAppend(t *testing.T) {
	assert := assert.New(t)
	t.Run("append an empty list", func(t *testing.T) {
		list := NewLinkedList[string]()

		valueToAppend := "first item"
		list.append(valueToAppend)

		listLength := list.GetLength()

		assert.Equal(listLength, 1)
		assert.Equal(list.Head.Value, valueToAppend)
		assert.Equal(list.Tail.Value, valueToAppend)

	})

	t.Run("append to a list with some values already in it", func(t *testing.T) {
		list := NewLinkedList[string]()

		firstValueToAppend := "value to append"
		secondValueToAppend := "second value to append"

		list.append(firstValueToAppend)
		list.append(secondValueToAppend)
		currentNode := list.Head.Next.Value

		assert.Equal(list.length, 2)
		assert.Equal(list.Head.Value, firstValueToAppend)
		assert.Equal(list.Tail.Value, secondValueToAppend)
		assert.Equal(currentNode, secondValueToAppend)
	})

	t.Run("prepend to a empty list", func(t *testing.T) {

		list := NewLinkedList[string]()

		valueToPrepend := "value to prepend"

		list.prepend(valueToPrepend)

		assert.Equal(list.length, 1)
		assert.Equal(list.Head.Value, valueToPrepend)
		assert.Equal(list.Tail.Value, valueToPrepend)
	})

	t.Run("prepend with a list with a least 1 item", func(t *testing.T) {
		list := NewLinkedList[string]()
		firstValueToPrepend := "first value"
		secondValueToPrepend := "second value"

		list.prepend(firstValueToPrepend)
		list.prepend(secondValueToPrepend)

		assert.Equal(list.length, 2)
		assert.Equal(list.Head.Value, secondValueToPrepend)
		assert.Equal(list.Tail.Value, firstValueToPrepend)
	})

}
