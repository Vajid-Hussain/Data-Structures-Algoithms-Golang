// func rangeSum(nums []int, n int, left int, right int) int {
//     var (
//         result []int
//         finalResult int
//     )

//     for i:=0; i<n; i++{
//         result= append(result, nums[i])
//         for j:=i+1; j<n; j++{
//             result=append(result, nums[j]+result[len(result)-1])
//         }
//     }
//     // fmt.Println(result)
//     slices.Sort(result) 
//     for i:=left-1; i<right; i++{
//         finalResult+= result[i]
//     }
//     return finalResult % 1000000007
// }