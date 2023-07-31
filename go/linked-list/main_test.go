package main

import (
	"fmt"
	"testing"
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
	t.Run("append an empty list", func(t *testing.T) {
		list := NewLinkedList[string]()

		valueToAppend := "first item"
		list.append(valueToAppend)

		listLength := list.GetLength()

		if listLength != 1 {
			testError(t, "listLength", 1, listLength)
		}

		if list.Head.Value != valueToAppend {
			testError(t, "Head.Value", valueToAppend, list.Head.Value)
		}

		if list.Tail.Value != valueToAppend {
			testError(t, "Tail.Value", valueToAppend, list.Tail.Value)
		}
	})

	t.Run("append to a list with some values already in it", func(t *testing.T) {
		list := NewLinkedList[string]()

		firstValueToAppend := "value to append"
		secondValueToAppend := "second value to append"

		list.append(firstValueToAppend)
		list.append(secondValueToAppend)

		if list.length != 2 {
			testError(t, "length", 2, list.length)
		}

		if list.Head.Value != firstValueToAppend {
			testError(t, "List.Head.Value", fmt.Sprintf("!= %s", firstValueToAppend), list.Head.Value)
		}

		if list.Tail.Value != secondValueToAppend {
			testError(t, "List.Tail.Value", secondValueToAppend, list.Tail.Value)
		}

		currentNode := list.Head.Next.Value

		if currentNode != secondValueToAppend {
			testError(t, "node.value", secondValueToAppend, currentNode)
		}
	})

	t.Run("prepend to a empty list", func(t *testing.T) {

		list := NewLinkedList[string]()

		valueToPrepend := "value to prepend"

		list.prepend(valueToPrepend)

		if list.length != 1 {
			testError(t, "list.length", 1, list.length)
		}

		if list.Head.Value != valueToPrepend {
			testError(t, "list.head.value", valueToPrepend, list.Head.Value)
		}

		if list.Tail.Value != valueToPrepend {
			testError(t, "list.tail.value", valueToPrepend, list.Tail.Value)
		}
	})

	t.Run("prepend with a list with a least 1 item", func(t *testing.T) {
		list := NewLinkedList[string]()
		firstValueToPrepend := "first value"
		secondValueToPrepend := "second value"

		list.prepend(firstValueToPrepend)
		list.prepend(secondValueToPrepend)

		if list.length != 2 {
			testError(t, "list.length", 2, list.length)
		}
		if list.Head.Value != secondValueToPrepend {
			testError(t, "list.head.value", secondValueToPrepend, list.Head.Value)
		}
		if list.Tail.Value != firstValueToPrepend {
			testError(t, "list.tail.value", firstValueToPrepend, list.Tail.Value)
		}
	})

}
