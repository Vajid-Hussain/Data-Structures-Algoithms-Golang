package main

import "fmt"

func main() {
	var arr = []int{4, 65, 0, 7, 9, 3}
	fmt.Println("==", mergeSort(arr))
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	first := mergeSort(arr[len(arr)/2:])
	second := mergeSort(arr[:len(arr)/2])
	return merge(first, second)
}

func merge(arr1, arr2 []int) []int {
	var final []int
	var i, j int
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			final = append(final, arr1[i])
			i++
		} else {
			final = append(final, arr2[j])
			j++
		}
	}

	for ; i < len(arr1); i++ {
		final = append(final, arr1[i])
	}
	for ; j < len(arr2); j++ {
		final = append(final, arr2[j])
	}
	return final
}
