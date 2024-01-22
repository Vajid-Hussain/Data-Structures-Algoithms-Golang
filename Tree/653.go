// /**
//  * Definition for a binary tree node.
//  * type TreeNode struct {
//  *     Val int
//  *     Left *TreeNode
//  *     Right *TreeNode
//  * }
//  */
//  func findTarget(root *TreeNode, k int) bool {
// 	return TwoSum(root, k, map[int]bool{})
//  }
 
//  func TwoSum(root *TreeNode, k int, mp map[int]bool) bool{
// 	 if root == nil{
// 		 return false
// 	 }
// 	 if mp[k-root.Val]{
// 		 return true
// 	 }
// 	 mp[root.Val]=true
// 	 return TwoSum(root.Left, k, mp) || TwoSum(root.Right, k, mp)
//  }