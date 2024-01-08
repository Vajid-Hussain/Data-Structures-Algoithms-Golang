// package main

// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
//  func mergeNodes(head *ListNode) *ListNode {
//     // head=head.Next
//     current:=head
//     point:=head
//     var heads = true

//     for current != nil{
//         if current.Val != 0{
//             point.Val+=current.Val
//         }else{
//             if heads{
//                 current=current.Next
//                 heads=false
//                 continue
//             }
//             point.Next=current.Next
//             point=point.Next
//             current=current.Next
//             if current==nil{
//                 continue
//             }
//         }
//         current=current.Next
//     }
//     return head
// }