package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

type start struct {
	head *Node
}

func (list *start) insert(data int) {

	var node Node
	node.Val = data
	node.Next = nil

	if list.head == nil {
		list.head = &node
		return
	}

	current := list.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &node
}

func main() {
	var head = new(start)

	head.insert(1)
	head.insert(2)

	current := head.head
	// fmt.Println("##", current)

	for current != nil {
		fmt.Println("!!", current.Val)
		current = current.Next
	}
}
