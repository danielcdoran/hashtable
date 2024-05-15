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
