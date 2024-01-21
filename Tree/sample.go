package main

import (
	"fmt"
	"math"
)

type node struct {
	Left  *node
	Val   int
	Right *node
}

type tree struct {
	Root *node
}

func (t *tree) insert(data int) {
	if t.Root == nil {
		t.Root = &node{Val: data}
	} else {
		t.Root.insert(data)
	}
}

func (t *node) insert(data int) {
	if t.Val < data {
		if t.Right == nil {
			t.Right = &node{Val: data}
		} else {
			t.Right.insert(data)
		}
	} else {
		if t.Left == nil {
			t.Left = &node{Val: data}
		} else {
			t.Left.insert(data)
		}
	}
}

func print(node *node) {
	if node == nil {
		return
	}
	print(node.Left)
	fmt.Println(node.Val)
	print(node.Right)
}

func validate(node *node, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min || node.Val >= max {
		return false
	}

	first := validate(node.Left, min, node.Val)
	second := validate(node.Right, node.Val, max)

	return first && second
}

func closest(node *node, target int) int {
	closestvar := node.Val

	pos := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for node != nil {

		if pos(node.Val-target) < pos(closestvar-target) {
			closestvar = node.Val
		}

		if node.Val < target {
			node = node.Right
		} else if node.Val > target {
			node = node.Left
		} else {
			return closestvar
		}
	}
	return closestvar
}

func delete(node *node, target int) *node {
	if node == nil {
		return nil
	} else if node.Val < target {
		node.Right = delete(node.Right, target)
	} else if node.Val > target {
		node.Left = delete(node.Left, target)
	} else {
		if node.Left == nil && node.Right == nil {
			return nil
		} else if node.Right == nil {
			return node.Left
		} else if node.Left == nil {
			return node.Right
		} else {
			node.Val = max(node.Left)
			node.Left = delete(node.Left, node.Val)
		}
	}
	return node
}

func max(node *node) int {
	for node.Right != nil {
		node = node.Right
	}
	return node.Val
}

func main() {
	tr := tree{}
	tr.insert(3)
	tr.insert(5)
	tr.insert(1)
	tr.insert(2)
	tr.insert(40)
	tr.Root = delete(tr.Root, 3)

	fmt.Println("closet value is", closest(tr.Root, 3))
	fmt.Println("valid:=", validate(tr.Root, math.MinInt, math.MaxInt))
	print(tr.Root)
}
