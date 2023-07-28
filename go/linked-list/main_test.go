package main

import (
	"testing"
)

func testError(t *testing.T, field string, expected any, got any) {
	t.Errorf("Field %s expected to be '%v' but got '%s'", field, expected, got)
}

func TestLinkedListAppend(t *testing.T) {
	t.Run("append an empty list", func(t *testing.T) {
		list := NewLinkedList[string]()

		valueToAppend := "first item"
		list.append(valueToAppend)

		listLength := list.getLength()

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
}
