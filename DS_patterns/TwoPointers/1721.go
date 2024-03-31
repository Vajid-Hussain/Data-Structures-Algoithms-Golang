// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
//  func swapNodes(head *ListNode, k int) *ListNode {
//     var left =head
//     for i:=1; i<k; i++{
//         left=left.Next
//     }

//     var right=head
//     var end = left
//     for end.Next!=nil{
//         right=right.Next
//         end=end.Next
//     }

//     left.Val, right.Val= right.Val, left.Val
//     return head
// }