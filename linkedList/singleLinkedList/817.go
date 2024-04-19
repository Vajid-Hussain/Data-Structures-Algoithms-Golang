// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
//  func numComponents(head *ListNode, nums []int) int {
// 	var (
// 		result int
// 		mp     = map[int]bool{}
// 	)

// 	for _, val := range nums {
// 		mp[val] = true
// 	}

// 	for head != nil {
// 		if mp[head.Val] {
//             // fmt.Println(head)
// 			for  head!=nil  && mp[head.Val]  {
// 				mp[head.Val] = false
// 				head = head.Next
// 			}
//             result++
// 		} else {
// 			head = head.Next
// 		}
// 	}

// 	// fmt.Println(mp, result)
// 	return result
// }