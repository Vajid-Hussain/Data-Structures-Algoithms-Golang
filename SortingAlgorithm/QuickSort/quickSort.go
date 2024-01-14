package main

import (
	"fmt"
)

// func Partition(arr []int, low int, high int) ([]int, int) {
// 	pivet := arr[high]
// 	i := low

// 	for j := low; j < high; j = j + 1 {
// 		if arr[j] < pivet {
// 			arr[i], arr[j] = arr[j], arr[i]
// 			i++
// 		}
// 	}
// 	arr[i], arr[high] = arr[high], arr[i]
// 	return arr, i
// }

// func quickSort(arr []int, low, high int) []int {
// 	if low < high {
// 		var p int
// 		arr, p = Partition(arr, low, high)
// 		arr = quickSort(arr, low, p-1)
// 		arr = quickSort(arr, p+1, high)
// 	}
// 	return arr
// }

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivet := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivet {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func main() {
	var arr = []int{5, 2, 8, 5, 6, 1, 9}
	fmt.Println(quickSort(arr, 0, len(arr)-1))
}
