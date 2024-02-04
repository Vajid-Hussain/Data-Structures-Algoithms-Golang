// package main

// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
//  func removeNthFromEnd(head *ListNode, n int) *ListNode {

//     if head.Next==nil && n==1{
//         return nil
//     }

//     current:=head
//     var length int

//     for current!=nil{
//         current=current.Next
//         length++
//     }

//     length-=(n)
//     i:=1
//     current=head.Next
//     previous:=head

//     if i>length{
//         // head.Next=nil
//         return head.Next
//     }

//     for i<=length{
//         fmt.Println(i,length)
//         if i==length{
//             previous.Next=current.Next
//         }

//         previous=current
//         current=current.Next
//         i++
//     }

//     return head
// }