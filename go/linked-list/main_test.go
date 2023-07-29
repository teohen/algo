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

		if list.Head == nil {
			testError(t, "Head", "not nil", "nil")
		}

		if list.Tail == nil {
			testError(t, "Tail", "not nil", "nil")
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

		valueToAppend := "value to append"

		list.append(valueToAppend)
		list.append(valueToAppend)

		if list.length != 2 {
			testError(t, "length", 2, list.length)
		}

		if list.Head.Value == valueToAppend {
			testError(t, "List.Head.Value", fmt.Sprintf("!= %s", valueToAppend), valueToAppend)
		}

		if list.Tail.Value != valueToAppend {
			testError(t, "List.Tail.Value", valueToAppend, list.Tail.Value)
		}

		currentNode := list.Head.Next.Value

		if currentNode != valueToAppend {
			testError(t, "node.value", valueToAppend, currentNode)
		}
	})

}
