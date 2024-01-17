// package main

// import "fmt"

// type LinkdeList struct {
// 	Val  int
// 	Next *LinkdeList
// }

// var Head *LinkdeList

// func delete(val int) {
// 	if val == Head.Val {
// 		pop()
// 		return
// 	}
// 	currentNodeData := Head.Val
// 	pop()
// 	delete(val)

// 	push(currentNodeData)
// }

// func push(data int) {
// 	node := &LinkdeList{}
// 	node.Val = data
// 	node.Next = Head
// 	Head = node
// }

// func pop() {
// 	Head = Head.Next
// }

// func print() {
// 	current := Head
// 	for current != nil {
// 		fmt.Println(current.Val)
// 		current = current.Next
// 	}
// }

// func main() {
// 	push(4)
// 	push(8)
// 	push(9)
// 	delete(8)
// 	delete(4)
// 	delete(9)
// 	push(2)
// 	print()
// }
