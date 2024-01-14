// package main

// import "fmt"

// func main() {
// 	arr := []int{4, 7, 1, 0, 9, 6}

// 	for i := 1; i < len(arr); i++ {
// 		for j := 0; j < (len(arr) - i); j++ {
// 			if arr[j] < arr[j+1] {
// 				arr[j] = arr[j] + arr[j+1]
// 				arr[j+1] = arr[j] - arr[j+1]
// 				arr[j] = arr[j] - arr[j+1]
// 			}
// 		}
// 	}
// 	fmt.Println("@@", arr)
// }
