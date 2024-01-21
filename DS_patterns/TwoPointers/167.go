// func twoSum(numbers []int, target int) []int {
//     start:=0
//     end:=len(numbers)-1
//     var currentSum int

//     for start < end{
//         currentSum=numbers[start]+ numbers[end]
//         if currentSum==target{
//             return []int{start+1, end+1}
//         }
//         if currentSum>target{
//             end-=1
//         }
//         if currentSum< target{
//             start+=1
//         }
//     }
//    return []int{}
// }