// package main

// import (
// 	"fmt"
// 	"time"
// )

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// type Start struct {
// 	Head *ListNode
// }

// func (head *Start) insertion(data int) {
// 	node := &ListNode{}
// 	node.Val = data

// 	if head.Head == nil {
// 		head.Head = node
// 		return
// 	}

// 	current := head.Head
// 	for current.Next != nil {
// 		current = current.Next
// 	}

// 	current.Next = node
// }

// func (head *Start) SelectionSort() {
// 	initial := head.Head

// 	for initial != nil {
// 		fmt.Println("##", initial)
// 		small := initial
// 		current := initial.Next
// 		var previous *ListNode
// 		previousNode := &ListNode{}

// 		temp := &ListNode{Next: initial}

// 		firstNodePrevious := temp

// 		for current != nil {
// 			// fmt.Println("!!", current)
// 			if small.Val > current.Val {
// 				previousNode = previous
// 				small = current
// 			}
// 			current = current.Next
// 			previous = temp.Next
// 		}

// 		if small != initial {
// 			if head.Head == initial {
// 				firstNodeNext := head.Head.Next
// 				secondNodeNext := small.Next
// 				firstNode := head.Head

// 				small.Next = firstNodeNext
// 				head.Head.Next = secondNodeNext

// 				head.Head = small
// 				previousNode.Next = firstNode
// 			} else {
// 				firstNodeNext := initial.Next
// 				SecondNodeNext := small.Next
// 				firstNode := initial

// 				initial.Next = SecondNodeNext
// 				small.Next = firstNodeNext

// 				firstNodePrevious.Next = small
// 				previous.Next = firstNode
// 			}
// 		}
// 		fmt.Println("**")
// 		loop := head.Head
// 		for loop != nil {
// 			fmt.Println(loop)
// 			loop = loop.Next
// 			time.Sleep(2 * time.Second)
// 		}
// 	}
// }

// func main() {
// 	var linkedList Start
// 	linkedList.insertion(4)
// 	linkedList.insertion(2)
// 	linkedList.insertion(1)
// 	linkedList.insertion(3)

// 	linkedList.SelectionSort()

// 	current := linkedList.Head

// 	for current != nil {
// 		fmt.Println(current.Val)
// 		current = current.Next
// 	}
// }
