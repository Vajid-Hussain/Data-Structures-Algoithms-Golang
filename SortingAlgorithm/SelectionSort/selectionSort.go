package main

import "fmt"

func main() {
	var arr = []int{5, 6, 1, 2, 0, 4, 7, 3}

	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if min != i {
			temp := arr[min]
			arr[min] = arr[i]
			arr[i] = temp
		}
	}

	fmt.Println("##", arr)
}
