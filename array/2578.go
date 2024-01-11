// package main

// func main() {
// 	var num int
// 	num = 4325
// 	var remainder int
// 	var arr []int
// 	var i int

// 	for num != 0 {
// 		remainder = num % 10
// 		arr = append(arr, remainder)
// 		i++
// 		num = num / 10
// 	}

// 	for i := 0; i <= len(arr)-2; i++ {
// 		small := i
// 		for j := i + 1; j <= len(arr)-1; j++ {
// 			if arr[small] > arr[j] {
// 				small = j
// 			}
// 		}
// 		if i != small {
// 			arr[i] = arr[i] + arr[small]
// 			arr[small] = arr[i] - arr[small]
// 			arr[i] = arr[i] - arr[small]
// 		}
// 	}
// 	var word1, word2 int
// 	for i := 0; i < len(arr); {
// 		word1 *= 10
// 		word1 += arr[i]
// 		i++

// 		if i == len(arr)-1 {
// 			break
// 		}

// 		word2 *= 10
// 		word2 += arr[i]
// 		i++
// 	}
// 	result := word1 + word2
// 	return result
// }
