// package main

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func rightSideView(root *TreeNode) []int {
// 	var rightNode []int
// 	current := root.Right
// 	for current != nil {
// 		rightNode = append(rightNode, current.Val)
// 		current = current.Right
// 	}
// 	return rightNode
// }