// /**
//  * Definition for a binary tree node.
//  * type TreeNode struct {
//  *     Val int
//  *     Left *TreeNode
//  *     Right *TreeNode
//  * }
//  */

//  var result []string

//  func binaryTreePaths(root *TreeNode) []string {
// 	 var currentPos string
// 	 childPath(root, currentPos)
// 	 current := result
 
// 	 result = nil
// 	 return current
//  }
 
//  func childPath(node *TreeNode, currentPos string) {
// 	 if node == nil {
// 		 return
// 	 }
// 	 if currentPos == "" {
// 		 currentPos += fmt.Sprint(node.Val)
// 	 } else {
// 		 currentPos += fmt.Sprint("->", node.Val)
// 	 }
// 	 fmt.Println(currentPos)
 
// 	 childPath(node.Left, currentPos)
// 	 childPath(node.Right, currentPos)
// 	 if node.Left == nil && node.Right == nil {
// 		 result = append(result, currentPos)
// 		 currentPos = ""
// 	 }
//  }