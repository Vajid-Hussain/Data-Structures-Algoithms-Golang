// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */

// // type keep struct{
// //     Next *ListNode
// // }
// func deleteDuplicates(head *ListNode) *ListNode {
//     var (
//         // mp= make(map[int]int)
//         prev =&ListNode{Next:head}
//     )

//     // for current!=nil{
//     //     mp[current.Val]++
//     //     current=current.Next
//     // }

//     result:= prev

//     for head.Next!=nil{

//         if head.Next.Val==head.Val{
//             for head.Next!=nil && head.Val == head.Next.Val{
//                 head=head.Next
//             }
//             prev= head.Next
//             prev=prev.Next
//         }else{
//             prev=prev.Next
//         }

//         // val,ok:=mp[current.Val]
//         // if ok{
//         //     if val>1{
//         //         prev.Next=prev.Next.Next
//         //     }else{
//         //         prev=prev.Next
//         //     }
//         // }
//         head=head.Next
//     }

//     return result.Next
// }