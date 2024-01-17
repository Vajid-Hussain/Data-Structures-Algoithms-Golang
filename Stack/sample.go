// package main

// import "fmt"

// var arr = []int{4, 5, 6, 3, 2, 8, 9}
// var top = 6

// func delete(arr []int, pos int) {
// 	if pos == arr[top] {
// 		return
// 	}

// 	data := arr[top]
// 	top--
// 	delete(arr, pos)
// 	arr[top] = data
// 	top++
// }

// func main() {

// 	delete(arr, 5)
// 	for i := 0; i < top; i++ {
// 		fmt.Println("@@", arr[i])
// 	}
// 	fmt.Println("**", arr)

// }
