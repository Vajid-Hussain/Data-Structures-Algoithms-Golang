package main

import "fmt"

func main() {
	arr := []int{4, 4, 2, 0, 0, 9, 4}

	for i := 1; i < len(arr); i++ {
		j := i

		for j > 0 && arr[j] < arr[j-1] {
			temp := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = temp
			j--
		}
	}
	fmt.Println("@@", arr)
}
