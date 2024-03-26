// func longestSubarray(nums []int) int {
// 	var result int
// 	var exist bool
// 	var count int
// 	var pos int

// 	for i := 0; i < len(nums); i++ {
// 		if !exist {
// 			count++
// 			if nums[i] == 0 {
// 				exist = true
// 				pos = i
// 			}
// 			continue
// 		}

// 		if nums[i] == 0 {
// 			if result < count {
// 				result = count - 1
// 			}
// 			count = i - pos
// 			pos = i
// 		} else {
// 			count++
// 		}
// 	}

// 	if result < count {
// 		result = count - 1
// 	}
// 	return result
// }