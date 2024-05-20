package hello

import (
	"testing"
)

func TestGiven1Item_then_Return(t *testing.T) {
	list := LinkedList{}
	list.insert(10, "string")
	actual, _ := list.get(10)
	if actual != "string" {
		t.Errorf("got %q, expected %q", actual, "string")
	}
}

func TestGivenZeroItems_then_ReturnError(t *testing.T) {
	list := LinkedList{}
	_, err := list.get(10)
	if err == nil {
		t.Errorf("got %q, expected nil", err)
	}
}

func TestGivenSameIndex_then_UpdateData(t *testing.T) {
	list := LinkedList{}
	list.insert(10, "string")
	list.insert(10, "nextstring")
	actual, _ := list.get(10)
	if actual != "nextstring" {
		t.Errorf("got %q, expected %q", actual, "nextstring")
	}
}

func TestGivenAddedInReverseOrder_then_DataReturned(t *testing.T) {
	list := LinkedList{}
	list.insert(12, "12string")
	list.insert(9, "9string")
	list.insert(3, "3string")
	actual, _ := list.get(3)
	if actual != "3string" {
		t.Errorf("got %q, expected %q", actual, "3string")
	}
}

func TestLargestIndex_then_AddToEnd(t *testing.T) {
	list := LinkedList{}
	list.insert(1, "1string")
	list.insert(9, "9string")
	list.insert(33, "33string")
	list.Display()
	actual, _ := list.get(33)
	if actual != "33string" {
		t.Errorf("got %q, expected %q", actual, "33string")
	}
}
