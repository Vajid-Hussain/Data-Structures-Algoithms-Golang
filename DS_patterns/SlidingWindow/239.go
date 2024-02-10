// // func maxSlidingWindow(nums []int, k int) []int {
// // 	var maxArr []int
// //     var visited =false
// // 	for i := k - 1; i < len(nums); i++ {
// // 		pos := i - (k - 1)
// // 		var max = nums[pos]
// //         if visited && maxArr[len(maxArr)-1] < nums[pos+(k-1)]  {
// //             maxArr=append(maxArr, nums[pos+(k-1)])
// //             continue
// //         }
// //         if visited && maxArr[len(maxArr)-1] > nums[pos-1]{
// //             maxArr=append(maxArr,maxArr[len(maxArr)-1])
// //             continue
// //         }
// //         visited = true
// // 		for j := pos; j < pos+k; j++ {
// // 			if max < nums[j] {
// // 				max = nums[j]
// // 			}
// // 		}
// // 		maxArr = append(maxArr, max)
// // 	}
// // 	return maxArr
// // }

// func maxSlidingWindow(nums []int, k int) []int {
//     deque := make([]int, 0, k)
//     result := make([]int, 0, len(nums) - k + 1)

//     for i, n := range nums {
//         // push - step
//         // remove all indeces from deque that corresponds
//         // to numbers less then number we are trying to push
//         // this will keep deque in ascending order
//         // so maximu would alwais be in deque[0] 
//         // for input: 1,3,-1,-3,5,3,6,7
//         // deque:
//         // [1]
//         // [3]
//         // [3 -1]
//         // [3 -1 - 3]
//         // [5]
//         // ...etc
//         for len(deque) > 0 && nums[deque[len(deque) - 1]] < n {
//             deque = deque[:len(deque) - 1]
//         }
//         deque = append(deque, i)

//         // check if we reach window size. N = 8, k = 3, i < 2 -> 0, 1, *2*
//         if i < k - 1 {
//             continue
//         }

//         // deque[0] is alwais maximum
//         result = append(result, nums[deque[0]])

//         // pop - step
//         // here we have to remove index of window start if it is in deque
//         // for example nums with winow: 1 [3  -1  -3] 5  3  6  7
//         // N = 8, k = 3, current i = 3 (-3 was pushed)
//         // we check if index i - k + 1 = 3 - 3 + 1 = 1
//         // is in deque at index 0
//         // index i - k + 1 could be only at deque[0]
//         // otherwise we removed it on push step.
//         // 
//         // if we add 3 to deque and its a max
//         // when we add next numbers there are 2 way
//         // 1. 3 is less new number -> new number will remove 3 on push step
//         // 2. 3 is greate all new numbers, so it alwais be at deque[0]
//         if deque[0] == i - k + 1{
//             deque = deque[1:]
//         }
//     }

//     return result

// }