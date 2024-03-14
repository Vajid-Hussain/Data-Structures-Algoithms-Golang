// // func numberOfSubarrays(nums []int, k int) int {
// // 	var (
// // 		result int
// // 		count  int
// // 		j      int
// // 	)

// // 	for i := 0; i < len(nums); i++ {
// // 		if nums[i]%2 != 0 {
// // 			count++
// // 		}
// // 		// fmt.Println(i, count)

// // 		if count == k {
// // 			// fmt.Println(i,"+++++")
// // 			result++
// // 		} else if count > k {
// // 			// fmt.Println("----------", i)
// // 			count = 0
// // 			i = j
// // 			j++
// // 		} else if (len(nums)-1) == i && count < k {
// // 			// fmt.Println("break")
// // 			break
// // 		}

// //         if (len(nums)-1) == i && count == k {
// // 			// fmt.Println("==========", i)
// // 			count = 0
// // 			i = j
// // 			j++
// // 		} 
// // 	}
// // 	return result
// // }

// func numberOfSubarrays(nums []int, k int)int{
//     var (
//         start int
//         end int
//         result int
//         oddCount int 
//         temp int 
//     )

//     for end <len(nums){
//         if nums[end] %2 !=0{
//             oddCount++
//             temp=0
//         }

//         for oddCount==k{
//             temp++
//             if nums[start] %2 !=0 {
//                 oddCount--
//             }
//             start++
//         }
//         result += temp
//         fmt.Println(result)
//         end++
//     }
//     return result
// }