// /**
//  * Definition for a binary tree node.
//  * type TreeNode struct {
//  *     Val int
//  *     Left *TreeNode
//  *     Right *TreeNode
//  * }
//  */
//  func bstFromPreorder(preorder []int) *TreeNode {
//     var head= &TreeNode{}

//     for i,val:=range preorder{
//         if i==0{
//             head=&TreeNode{Val:val}
//             continue
//         }
//         insert(val, head)
//     }
//     return head
// }

// func insert(val int, node *TreeNode) {
//     if val>node.Val{
//         if node.Right==nil{
//             node.Right=&TreeNode{Val:val}
//         }else{
//             insert(val, node.Right)
//         }
//     }else{
//         if node.Left ==nil{
//             node.Left= &TreeNode{Val:val}
//         }else{
//             insert(val, node.Left)
//         }
//     }
// }
