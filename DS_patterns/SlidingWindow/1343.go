// func numOfSubarrays(arr []int, k int, threshold int) int {
// 	var (
// 		result  int
// 		start   int
// 		current int
// 		total   int
// 	)

// 	for i := 0; i < k-1; i++ {
// 		total += arr[i]
// 		current++
// 	}

// 	for current < len(arr) {
//         total+=arr[current]

// 		if total/k >= threshold {
// 			result++
// 		}
// 		total -= arr[start]
// 		start++

// 		current++
// 	}

// 	return result
// }