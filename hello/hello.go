package hello

import (
	"errors"
	"fmt"
)

// User user type
type User struct {
	ID   int64
	Name string
	Addr *Address
}

// Address address type
type Address struct {
	City   string
	ZIP    int
	LatLng [2]float64
}
type Node struct {
	index int
	data  string
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Display() {
	current := list.head

	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}

	fmt.Print("Linked list: ")
	for current != nil {
		fmt.Printf("%d %s ", current.index, current.data)
		current = current.next
	}
	fmt.Println()
}

func (list *LinkedList) insert(index int, data string) {
	newNode := &Node{index: index, data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}
	var prevNode *Node
	current := list.head
	for current.next != nil {
		if current.index == index {
			current.data = data
			break
		}
		if index < current.index {
			newNode.next = current
			if prevNode == nil {
				list.head = newNode
			} else {
				prevNode.next = current
			}
			return
		}
		prevNode = current
		current = current.next
	}

	current.next = newNode
}

// Assumes items are ordered in bucket by increasing index
func (list LinkedList) get(index int) (string, error) {

	if list.head == nil {
		return "", errors.New("not found")
	}

	current := list.head
	for {
		if current.index == index {
			return current.data, nil
		}
		if current.next == nil {
			break
		}
		current = current.next
	}
	return "", errors.New("not found")
}

var alex = User{}

// Hello writes a welcome string
func Hello() string {
	return "Hello, " + alex.Name
}
