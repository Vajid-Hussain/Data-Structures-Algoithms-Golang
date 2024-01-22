// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// type node struct {
// 	Left  *node
// 	Val   int
// 	Right *node
// }

// type root struct {
// 	root *node
// }

// func (r *root) inset(data int) {
// 	if r.root == nil {
// 		r.root = &node{Val: data}
// 	} else {
// 		r.root.insert(data)
// 	}
// }

// func (t *node) insert(data int) {
// 	if data < t.Val {
// 		if t.Left == nil {
// 			t.Left = &node{Val: data}
// 		} else {
// 			t.Left.insert(data)
// 		}
// 	}
// 	if data > t.Val {
// 		if t.Right == nil {
// 			t.Right = &node{Val: data}
// 		} else {
// 			t.Right.insert(data)
// 		}
// 	}
// }

// func print(w io.Writer, node *node, ns int, ch rune) {
// 	if node == nil {
// 		return
// 	}
// 	for i := 0; i < ns; i++ {
// 		fmt.Fprint(w, " ")
// 	}
// 	fmt.Fprintf(w, "%c:%v\n", ch, node.Val)
// 	print(w, node.Left, ns+2, 'L')
// 	print(w, node.Right, ns+2, 'R')
// }

// func main() {
// 	tree := &root{}
// 	tree.inset(8)
// 	tree.inset(9)
// 	tree.inset(15)
// 	tree.inset(2)
// 	print(os.Stdout, tree.root, 0, 'M')
// }
