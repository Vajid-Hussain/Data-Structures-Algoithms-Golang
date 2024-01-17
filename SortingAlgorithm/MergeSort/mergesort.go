// package main

// import "fmt"

// func mergeSort(arr []int) []int {
// 	if len(arr) < 2 {
// 		return arr
// 	}
// 	first := mergeSort(arr[:len(arr)/2])
// 	second := mergeSort(arr[len(arr)/2:])
// 	return merge(first, second)
// }

// func merge(arr1, arr2 []int) []int {
// 	final := []int{}
// 	i := 0
// 	j := 0

// 	for i < len(arr1) && j < len(arr2) {
// 		if arr1[i] < arr2[j] {
// 			final = append(final, arr1[i])
// 			i++
// 		} else {
// 			final = append(final, arr2[j])
// 			j++
// 		}
// 	}
// 	for ; i < len(arr1); i++ {
// 		final = append(final, arr1[i])
// 	}
// 	for ; j < len(arr2); j++ {
// 		final = append(final, arr2[j])
// 	}
// 	return final
// }

// func main() {
// 	arr := []int{5, 3, 7, 1, 8, 3, 0}
// 	result := mergeSort(arr)
// 	fmt.Println(result)
// }
