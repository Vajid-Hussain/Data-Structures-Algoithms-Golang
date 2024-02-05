// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
//  func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
//     current:=&ListNode{Next:list1}

//     for i:=1; i<=a; i++{
//         current=current.Next
//     }
//     current.Next, current= list2, current.Next

//     for i:=a; i<b; i++{
//         current=current.Next
//     }

//     newList:=list1
//     for newList.Next!=nil{
//         newList=newList.Next
//     }
//     newList.Next=current.Next

//     return list1
// }