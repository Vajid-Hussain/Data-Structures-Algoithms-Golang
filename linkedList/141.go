package linkedList141

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

type start struct {
	head *Node
}

func (list *start) insert(data int) {

	var node Node
	node.Val = data
	node.Next = nil

	if list.head == nil {
		list.head = &node
		return
	}

	current := list.head
	for current.Next != nil {
		current = current.Next
	}
	node.Next = list.head
	current.Next = &node
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// var mp = make(map[*ListNode]int)

	// for head != nil {
	// 	if _, ok := mp[head.Next]; ok {
	// 		return true
	// 	}
	// 	mp[head.Next]++
	// 	head = head.Next
	// }
	// return false

	slowPoint, fastPoint := head, head

	for fastPoint != nil && fastPoint.Next != nil {
		slowPoint = slowPoint.Next
		fastPoint = fastPoint.Next.Next
		if slowPoint == fastPoint {
			return true
		}
	}
	return false

}

func main() {
	var head = new(start)

	head.insert(1)
	head.insert(2)

	// current := head.head
	// val := hasCycle(head.head)
	// fmt.Println("**", val)

	// for current != nil {
	// 	fmt.Println("!!", current.Val)
	// 	current = current.Next
	// }
}
