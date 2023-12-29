// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
//  func reorderList(head *ListNode)  {
//     var (
//         fast = head
//         slow = head

//     )

//     for fast!=nil && fast.Next!=nil{
//         fast=fast.Next.Next
//         if fast==nil{
//             temp:=slow
//             slow=slow.Next
//             temp.Next=nil
//             break
//         }
//         slow=slow.Next
//     }

//     if fast!=nil{
//         temp:=slow
//         slow=slow.Next
//         temp.Next=nil
//     }

//     current:= slow
//     var prev *ListNode
//     for current!=nil{
//         next:= current.Next
//         current.Next=prev
//         prev=current
//         current=next
//     }

//     end:=prev

//     current=head
//     for current!=nil{
//         next:=current.Next
//         current.Next=end
//         if end!=nil{
//             end=end.Next
//         }
//         current=current.Next
//         if current!=nil{
//             current.Next=next
//             current=current.Next
//         }
//     }
 
// }