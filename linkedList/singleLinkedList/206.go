// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
//  var newHead *ListNode

//  func reverseList(head *ListNode) *ListNode {
// 	 if head==nil{
// 		 return head
// 	 }
// 	 recursions(head)
// 	 head.Next = nil
// 	 return newHead
//  }
 
//  func recursions(current *ListNode) {
// 	 if current.Next == nil {
// 		 newHead = current
// 		 return
// 	 }
// 	 prev := current
// 	 Next := current.Next
 
// 	 recursions(Next)
// 	 Next.Next = prev
//  }
 
//second solution

// func reverseList(head *ListNode) (prev *ListNode ){
//     prev, curt := nil, head
//     for curt!=nil{
//         nxt:=curt.Next
//         curt.Next=prev
//         prev=curt
//         curt=nxt
//     }
//     return prev
// }
