// /**
//  * Definition for a binary tree node.
//  * type TreeNode struct {
//  *     Val int
//  *     Left *TreeNode
//  *     Right *TreeNode
//  * }
//  */
//  func bstToGst(root *TreeNode) *TreeNode {
//     result:= root
//     var totalSum int

//     var sum func(*TreeNode)

//     sum= func(node *TreeNode) {
//         if node==nil{
//             return 
//         }

//         sum(node.Right)
//         totalSum+=node.Val
//         node.Val=totalSum
//         sum(node.Left)
//     }
//     sum(root)
//     return result
// }
