// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
//  func reverseEvenLengthGroups(head *ListNode) *ListNode {
//     var(
//         i =1
//         j =1
//         prev *ListNode
//         result = head
//     )

//     for head.Next!=nil{
//         j=1
//         prev= head
//         // fmt.Print("new ")

//         for j<=i && head.Next != nil && i!=1{
//             // fmt.Print(i)
//             head=head.Next
//             j++
//         }
//         fmt.Print(" ",head.Val, prev.Val, j-1, "  ")
//         if i!=1{
//             j--
//             prev= prev.Next
//         }

//         if j%2==0{
//             // fmt.Println( "recursion ", i, head.Val)
//             reverse(prev, j/2)
//         }

//         i++
//     }
//     // current:=result
//     // for current!=nil{
//     //     fmt.Println(current.Val)
//     //     current=current.Next
//     // }
//     return result
// }

// func reverse(head *ListNode, count int)*ListNode {
//     if count==0 {
//         return head
//     }
//     next:=reverse(head.Next, count-1)
//     head.Val, next.Val= next.Val, head.Val
//     return next.Next
// }