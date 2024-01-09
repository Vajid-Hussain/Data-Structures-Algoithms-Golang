package main

import "fmt"

type LinkedList struct {
	Val  int
	Next *LinkedList
}

type Start struct {
	Head *LinkedList
}

func (list *Start) Insertion(data int) {
	var node = &LinkedList{}

	node.Val = data
	node.Next = nil

	if list.Head == nil {
		list.Head = node
		return
	}

	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node

}

func main() {
	linkedList := new(Start)

	linkedList.Insertion(4)
	linkedList.Insertion(8)

	current := linkedList.Head

	for current != nil {
		fmt.Println("**", current.Val)
		current = current.Next
	}
}
