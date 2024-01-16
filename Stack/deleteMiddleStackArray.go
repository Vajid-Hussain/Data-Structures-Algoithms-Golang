// package main

// import "fmt"

// // var arr = []int{4, 5, 6, 3, 2, 8, 9}
// var arr = make([]int, 7)
// var top = 0
// var temp =true

// func delete(pos int) {
// 	if pos == arr[top] {
// 		pop()
// 		return
// 	}
// 	currentTop := arr[top]
// 	pop()

// 	delete(pos)
// 	push(currentTop)
// }

// func push(data int) {
// 	if temp{
// 		arr[top]=data
// 		temp=false
// 		return
// 	}
// 	top++
// 	arr[top] = data
// }

// func pop() {
// 	top = top -1
// }
	
// func main() {
// 	push(3)
// 	push(2)
// 	push(7)
// 	push(6)

// 	delete(6)
// 	for i := 0; i <= top; i++ {
// 		fmt.Println("@@", arr[i])
// 	}
// 	fmt.Println("**", arr)

// }
