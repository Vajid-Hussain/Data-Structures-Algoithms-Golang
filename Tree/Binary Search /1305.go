// /**
//  * Definition for a binary tree node.
//  * type TreeNode struct {
//  *     Val int
//  *     Left *TreeNode
//  *     Right *TreeNode
//  * }
//  */
//  func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
//     var (
//         result []int
//         mu sync.Mutex
//         ch =make(chan int)
//     )
//     go func(){
//         recursion(root1,&mu, ch)
//         recursion(root2,&mu, ch)
//         defer close(ch)
//     }()

//     for val:=range ch{
//         result=append(result, val)
//     }

//     // time.Sleep(time.Second*1)

//     slices.Sort(result)
//     // fmt.Println(result)
       
//     return result
// }

// func recursion(node *TreeNode, mu *sync.Mutex,ch chan int) {
//     if node==nil{
//         return
//     }

//     mu.Lock()
//     ch <- node.Val
//     mu.Unlock()
//     // fmt.Println("__",arr)

//     recursion(node.Left, mu, ch)
//     recursion(node.Right,  mu,ch)
// }
// //