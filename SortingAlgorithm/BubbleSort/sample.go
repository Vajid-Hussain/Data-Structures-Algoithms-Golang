package main

import "fmt"

func main() {
	arr := []int{3, 6, 1, 6, 5, 1}

	for n := 0; n < len(arr)-1; n++ {
		for i := 0; i < (len(arr)-n)-1; i++ {
			if arr[i] < arr[i+1] {
				arr[i] = arr[i] + arr[i+1]
				arr[i+1] = arr[i] - arr[i+1]
				arr[i] = arr[i] - arr[i+1]
			}
		}
	}
	fmt.Println("@@", arr)
}
