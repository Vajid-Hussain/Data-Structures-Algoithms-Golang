package main

import "fmt"

type Node struct {
	Prev *Node
	Val  int
	Next *Node
}

type Start struct {
	Head *Node
}

func (list *Start) Insertion(data int) {
	node := &Node{}
	node.Val = data

	if list.Head == nil {
		list.Head = node
		return
	}

	current := list.Head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = node
	node.Prev = current
}

func (list *Start) InsertBetween(data, pos int) {
	node := new(Node)
	node.Val = data
	i := 2

	current := list.Head
	for current != nil && i <= pos {
		if i == pos {
			node.Next = current.Next
			node.Prev = current
			current.Next = node
			return
		}
		current = current.Next
		i++
	}
}

func (list *Start) InsertFirst(data int) {
	node := new(Node)
	node.Val = data
	node.Next = list.Head
	list.Head.Prev = node
	list.Head = node
}

func (list *Start) DeleteEnd() {
	current := list.Head
	for current.Next.Next != nil {
		current = current.Next
	}
	current.Next = nil
}

func (list *Start) DeleteBigining() {
	list.Head = list.Head.Next
	list.Head.Prev = nil
}

func (list *Start) DeleteMid(pos int) {
	current := list.Head
	i := 1

	for current != nil {
		if i == pos {
			current.Prev.Next = current.Next
			current.Next.Prev = current.Prev
		}
		current = current.Next
		i++
	}
}
func main() {
	linkedList := Start{}
	linkedList.Insertion(4)
	linkedList.Insertion(5)
	linkedList.Insertion(7)
	linkedList.Insertion(8)
	linkedList.Insertion(9)
	linkedList.InsertBetween(6, 3)
	linkedList.InsertFirst(3)
	linkedList.DeleteEnd()
	linkedList.DeleteBigining()
	linkedList.DeleteMid(3)

	current := linkedList.Head

	for current != nil {
		fmt.Println("**", current)
		current = current.Next
	}
}
