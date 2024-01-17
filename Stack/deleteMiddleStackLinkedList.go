// package main

// import "fmt"

// type LinkdeList struct {
// 	Val  int
// 	Next *LinkdeList
// }

// var Head *LinkdeList

// func DeletePerticular(val int) {
// 	if val == Head.Val {
// 		Pop()
// 		return
// 	}
// 	currentTop := Head.Val
// 	Pop()

// 	DeletePerticular(val)
// 	Create(currentTop)
// }

// func Pop() {
// 	Head = Head.Next
// }

// func Create(data int) {
// 	node := &LinkdeList{}
// 	node.Val = data
// 	node.Next = Head
// 	Head = node
// }

// func Print() {
// 	current := Head

// 	for current != nil {
// 		fmt.Println("@@", current)
// 		current = current.Next
// 	}
// }

// func main() {
// 	Create(2)
// 	Create(8)
// 	Create(5)
// 	DeletePerticular(8)
// 	Print()
// }
