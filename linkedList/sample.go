// package main

// import "fmt"

// type LinkedList struct {
// 	Val  int
// 	Next *LinkedList
// }

// type Start struct {
// 	Head *LinkedList
// }

// func (list *Start) Insertion(data int) {
// 	var node = &LinkedList{}

// 	node.Val = data
// 	node.Next = nil

// 	if list.Head == nil {
// 		list.Head = node
// 		return
// 	}

// 	current := list.Head
// 	for current.Next != nil {
// 		current = current.Next
// 	}
// 	current.Next = node

// }

// func (list *Start) AddFirst(data int) {
// 	node := &LinkedList{}
// 	node.Val = data
// 	node.Next = list.Head

// 	list.Head = node
// }

// func (list *Start) AddBetween(data int, pos int) {
// 	i := 2
// 	current := list.Head
// 	var node = &LinkedList{}
// 	node.Val = data

// 	for i <= pos && current != nil {
// 		if i == pos {
// 			node.Next = current.Next
// 			current.Next = node
// 		}
// 		current = current.Next
// 		i++
// 	}
// }

// func main() {
// 	linkedList := new(Start)

// 	linkedList.Insertion(4)
// 	linkedList.Insertion(8)
// 	linkedList.Insertion(1)
// 	linkedList.AddFirst(2)
// 	linkedList.AddBetween(9, 3)

// 	current := linkedList.Head

// 	for current != nil {
// 		fmt.Println("**", current.Val)
// 		current = current.Next
// 	}
// }
