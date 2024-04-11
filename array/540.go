// func singleNonDuplicate(nums []int) int {
// 	var (
// 		left  int
// 		right = len(nums) - 1
// 		mid   int
// 	)

// 	if len(nums) == 1 {
// 		return nums[0]
// 	}

// 	for left <= right {
// 		fmt.Println(mid, "mid")
// 		mid = (left + right) / 2

// 		if mid == len(nums)-1 && nums[mid] != nums[mid-1] {
// 			return nums[mid]
// 		} else if mid == 0 && nums[mid] != nums[mid+1] {
// 			return nums[mid]
// 		} else if (nums[mid] != nums[mid+1]) && nums[mid] != nums[mid-1] {
// 			return nums[mid]
// 		} else {
// 			left, right = leftAndRight(left, right, mid, nums)
// 		}

// 	}
// 	return 0
// }

// func leftAndRight(left, right, mid int, nums []int) (int, int) {
// 	if mid%2 == 0 {
// 		if nums[mid] != nums[mid+1] {
// 			return left, mid - 1
// 		} else {
// 			return mid + 1, right
// 		}
// 	} else {
// 		if nums[mid] != nums[mid-1] {
// 			return left, mid - 1
// 		} else {
// 			return mid + 1, right
// 		}
// 	}
// 	return left, right
// }