// /**
//  * Definition for a binary tree node.
//  * type TreeNode struct {
//  *     Val int
//  *     Left *TreeNode
//  *     Right *TreeNode
//  * }
//  */
//  func deepestLeavesSum(root *TreeNode) int {
//     if root == nil {return 0}
//     var ans int
//     queue := []*TreeNode{root}

//     for len(queue) > 0{
//         nodeInCurrentLeavel := len(queue)
//         curLeavel:=0
//         for i:=0; i<nodeInCurrentLeavel; i++{
//             node:=queue[0]
//             queue= queue[1:]
//             curLeavel+=node.Val

//             if node.Left!=nil{
//                 queue =append(queue, node.Left)
//             }
//             if node.Right !=nil{
//                 queue = append(queue, node.Right)
//             }
//         }
//         ans=curLeavel
//     }
//     return ans
// }   