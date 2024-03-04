// func generateParenthesis(n int) []string {
//     result:= []string{}
//     current:=make([]byte, n*2)
//     rec(&result, n, 0,0, current)
//     return result
// }

// func rec(result *[]string, n int, left int , right int, current []byte){
//     if right+left==n*2{
//         *result=append(*result, string(current))
//         return
//     }

//     if left<n{
//         current[left+right] = '('
//         rec(result, n, left+1, right, current)
//     }
//     if right <left{
//         current[left+right]=')'
//         rec(result, n, left, right+1, current)
//     }
// }