// /**
//  * Definition for a binary tree node.
//  * type TreeNode struct {
//  *     Val int
//  *     Left *TreeNode
//  *     Right *TreeNode
//  * }
//  */
//  var result int

// func FindSum(root *TreeNode, low, high int){
//     if root == nil{
//         return 
//     }
//     if root.Val>=low && root.Val<=high{
//         result+=root.Val
//     }
//     FindSum(root.Right,low, high)
//     FindSum(root.Left, low, high)
// }

// func rangeSumBST(root *TreeNode, low int, high int) int {
//     result=0
//     FindSum(root, low, high)
//     return result
// }