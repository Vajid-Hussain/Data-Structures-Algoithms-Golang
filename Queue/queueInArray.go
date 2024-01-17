// package main

// import "fmt"

// func enque(queue *[]int, data int) {
// 	*queue = append(*queue, data)
// }

// func deque(queue *[]int) {
// 	*queue = (*queue)[1:]
// }

// func display(queue *[]int) {
// 	for _, val := range *queue {
// 		fmt.Println("@@", val)
// 	}
// }

// func main() {
// 	var queue []int
// 	enque(&queue, 4)
// 	enque(&queue, 5)
// 	// deque(&queue)
// 	display(&queue)
// }
