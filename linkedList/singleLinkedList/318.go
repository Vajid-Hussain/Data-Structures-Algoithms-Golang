/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func oddEvenList(head *ListNode) *ListNode {
    if head==nil{ return head}

    prev:=head
    secondNode:=head.Next
    current := head.Next

    for current!=nil && current.Next!=nil{
        prev.Next=current.Next
        prev=prev.Next
        current.Next=current.Next.Next
        current=current.Next
    }

    prev.Next=secondNode
    return head
}
