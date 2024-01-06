// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
//  func deleteMiddle(head *ListNode) *ListNode {
// 	var (
// 		slow = head
// 		fast = head
//         prevSlow *ListNode
// 	)

//     for fast!=nil && fast.Next!=nil{
//         prevSlow = slow
//         slow=slow.Next
//         fast=fast.Next.Next
//     }

//     if prevSlow==nil{
//         return prevSlow
//     }

//     prevSlow.Next=prevSlow.Next.Next
	
// 	return head
// }