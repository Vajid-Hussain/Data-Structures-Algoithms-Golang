// package main

// import (
// 	"fmt"
// 	"math"
// )

// type node struct {
// 	Left  *node
// 	Val   int
// 	Right *node
// }

// type Tree struct {
// 	Root *node
// }

// func (r *Tree) insert(data int) {
// 	if r.Root == nil {
// 		r.Root = &node{Val: data}
// 	} else {
// 		r.Root.insert(data)
// 	}
// }

// func (r *node) insert(data int) {
// 	if r.Val < data {
// 		if r.Right == nil {
// 			r.Right = &node{Val: data}
// 			return
// 		}
// 		r.Right.insert(data)
// 	} else {
// 		if r.Left == nil {
// 			r.Left = &node{Val: data}
// 			return
// 		}
// 		r.Left.insert(data)
// 	}
// }

// func print(root *node) {
// 	if root == nil {
// 		return
// 	}
// 	print(root.Left)
// 	fmt.Println(root.Val)
// 	print(root.Right)
// }

// func closest(node *node, target int) int {
// 	closet := node.Val

// 	dif := func(val int) int {
// 		if val < 0 {
// 			return -val
// 		}
// 		return val
// 	}

// 	for node != nil {
// 		if dif(node.Val-target) < dif(closet-target) {
// 			closet = node.Val
// 		}

// 		if node.Val > target {
// 			node = node.Left
// 		} else {
// 			node = node.Right
// 		}
// 	}
// 	return closet
// }

// func delete(node *node, target int) *node {
// 	if node == nil {
// 		return nil
// 	} else if node.Val < target {
// 		node.Right = delete(node.Right, target)
// 	} else if node.Val > target {
// 		node.Left = delete(node.Left, target)
// 	} else {
// 		if node.Left == nil && node.Right == nil {
// 			return nil
// 		} else if node.Left == nil {
// 			return node.Right
// 		} else if node.Right == nil {
// 			return node.Left
// 		} else {
// 			node.Val = maxLeft(node.Left)
// 			node.Left = delete(node.Left, node.Val)
// 		}
// 	}
// 	return node
// }

// func maxLeft(node *node) int {
// 	for node != nil && node.Right != nil {
// 		node = node.Right
// 	}
// 	return node.Val
// }

// func isValid(node *node, min, max int) bool {
// 	if node == nil {
// 		return true
// 	}
// 	if node.Val <= min || node.Val >= max {
// 		return false
// 	}

// 	first := isValid(node.Left, min, node.Val)
// 	second := isValid(node.Right, node.Val, max)
// 	return first && second
// }

// func main() {
// 	t := Tree{}
// 	t.insert(6)
// 	t.insert(9)
// 	t.insert(8)
// 	t.insert(10)
// 	// t.Root = delete(t.Root, 9)
// 	fmt.Println("##", isValid(t.Root, math.MinInt, math.MaxInt))
// 	fmt.Println("@@", closest(t.Root, 3))
// 	print(t.Root)
// }
