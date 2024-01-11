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
// 	i := 1
// 	var node = &LinkedList{}
// 	node.Val = data
// 	current := list.Head
// 	for current != nil {
// 		if i == pos {
// 			node.Next = current.Next
// 			current.Next = node
// 			return
// 		}
// 		current = current.Next
// 		i++
// 	}
// }

// func (list *Start) DeleteFirst() {
// 	list.Head = list.Head.Next
// }

// func (list *Start) DeleteEnd() {
// 	previous := list.Head
// 	current := list.Head.Next

// 	for current.Next != nil {
// 		previous = current
// 		current = current.Next
// 	}
// 	previous.Next = nil
// }

// func (list *Start) DeleteMid(pos int) {
// 	current := list.Head.Next
// 	previous := list.Head
// 	i := 1

// 	for current != nil {
// 		if pos == i {
// 			previous.Next = current.Next
// 			return
// 		}
// 		previous = current
// 		current = current.Next
// 		i++
// 	}
// }

// func main() {
// 	linkedList := new(Start)

// 	linkedList.Insertion(4)
// 	linkedList.Insertion(5)
// 	linkedList.Insertion(7)
// 	linkedList.AddFirst(2)
// 	linkedList.AddBetween(6, 3)
// 	linkedList.DeleteFirst()
// 	linkedList.DeleteEnd()
// 	linkedList.DeleteMid(1)

// 	current := linkedList.Head

// 	for current != nil {
// 		fmt.Println("**", current.Val)
// 		current = current.Next
// 	}
// }
