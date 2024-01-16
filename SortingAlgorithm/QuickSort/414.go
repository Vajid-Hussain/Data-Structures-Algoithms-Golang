// func thirdMax(nums []int) int {
//     if len(nums) == 0 { return -1 }
//     if len(nums) == 1 { return nums[0] }

//     // Building max heap
//     for i := len(nums)/2 - 1; i > -1; i-- {
//         heapDown(nums, i, 0, len(nums) - 1)
//     }

//     largest, res, count, low, high := nums[0], nums[0], 1, 0, len(nums) - 1

//     for low <= high {
//         if nums[low] != largest {
//             count++
//             largest = nums[low]
//         }

//         if count == 3 {
//             return largest
//         }

//         nums[low], nums[high] = nums[high], nums[low]
//         high--
//         heapDown(nums, 0, low, high)
//     }

//     return res
// }

// func heapDown(nums []int, position int, low int, high int) {
//     left := 2*position + 1
//     right := 2*position + 2
//     larger := position

//     if low <= left && left <= high {
//         if nums[left] > nums[larger] {
//             larger = left
//         }
//     }

//     if low <= right && right <= high {
//         if nums[right] > nums[larger] {
//             larger = right
//         }
//     }

//     if larger != position {
//         nums[larger], nums[position] = nums[position], nums[larger]
//         heapDown(nums, larger, low, high)
//     }
// } 