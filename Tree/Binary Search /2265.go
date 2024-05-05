// /**
//  * Definition for a binary tree node.
//  * type TreeNode struct {
//  *     Val int
//  *     Left *TreeNode
//  *     Right *TreeNode
//  * }
//  */
//  func averageOfSubtree(root *TreeNode) int {
//     var (
//         result int
//         postFunc func(*TreeNode)[2]int
//     )

//     postFunc = func(node *TreeNode) [2]int{
//         if node==nil{
//             return [2]int{0,0}
//         }

//         leftRes:= postFunc(node.Left)
//         rightRes:= postFunc(node.Right) 
//         computeAll:= [2]int{leftRes[0]+ rightRes[0]+ node.Val, leftRes[1]+ rightRes[1]+ 1}
//         // fmt.Println(computeAll)
//         if computeAll[0]/computeAll[1] == node.Val{
//             result++
//         }
//         return computeAll
//     }
//     postFunc(root)

//     return result
// }