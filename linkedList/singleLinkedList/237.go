// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
//  func deleteNode(node *ListNode) {
//     prev:=node
//     node=node.Next
//     for node.Next!=nil{
//         prev.Val=node.Val
//         prev=node
//         node=node.Next
//     }
//     prev.Val=node.Val
//     prev.Next=nil
// }