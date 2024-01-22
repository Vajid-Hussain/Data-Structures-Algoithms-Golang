// /**
//  * Definition for a binary tree node.
//  * type TreeNode struct {
//  *     Val int
//  *     Left *TreeNode
//  *     Right *TreeNode
//  * }
//  */
// func isValidBST(root *TreeNode) bool {
//     if root==nil{
//         return true
//     }
//     return check_valid(root, math.MinInt, math.MaxInt)
// }

// func check_valid(node *TreeNode, min,max int)bool{
//     if node==nil {
//         return true
//     }
//     if node.Val <=min || node.Val >=max{
//         return false
//     }

//     first:=check_valid(node.Right, node.Val,max)
//     Second:=check_valid(node.Left, min, node.Val)
//     return first && Second
// }