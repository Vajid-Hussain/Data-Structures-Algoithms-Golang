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

// By using gc for clean memory

// func twoSum(numbers []int, target int) []int {
// 	debug.SetGCPercent(1)

// var(
// 	i, j= 0, len(numbers)-1
// 	num int
// )

// for i<j{
// 	num = numbers[i] + numbers[j]
// 	if num < target{
// 		i++
// 	}else if num > target{
// 		j--
// 	}else{
// 		return []int{i+1, j+1}
// 	}
// }
// return []int{}
// }
