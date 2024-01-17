package main

import "fmt"

type LinkdeList struct {
	Val  int
	Next *LinkdeList
}

var Head *LinkdeList
var Tail *LinkdeList

func enque(data int) {
	node := &LinkdeList{}
	node.Val = data
	if Head == nil && Tail == nil {
		Head = node
		Tail = node
		return
	}
	Tail.Next = node
	Tail = node
}

func deque() {
	Head = Head.Next
}

func print() {
	current := Head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

func main() {
	enque(8)
	enque(7)
	enque(6)
	enque(5)
	print()
}
