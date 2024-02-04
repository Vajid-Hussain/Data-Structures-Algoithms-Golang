// package main
// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */

//  func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//     sum := 0
//     carry := 0
//     l3 := &ListNode{}
//     head := l3

//     for l1 != nil || l2 != nil {
//         sum = 0

//         if l1 != nil {
//             sum += l1.Val
//             l1 = l1.Next
//         }

//         if l2 != nil {
//             sum += l2.Val
//             l2 = l2.Next
//         }
//         sum += carry
//         carry = sum / 10
//         sum %= 10
//         l3.Next = &ListNode{Val: sum}
//         l3 = l3.Next
//     }

//     if carry != 0 {
//         l3.Next = &ListNode{Val: carry}
//     }

//     return head.Next
// }