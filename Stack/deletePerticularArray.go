// package main

// import "fmt"

// func push(stack *[]int, data int) {
// 	*stack = append((*stack), data)
// }

// func pop(stack *[]int) {
// 	*stack = (*stack)[:len(*stack)-1]
// }

// func delete(stack *[]int, val int) {
// 	if val == (*stack)[len(*stack)-1] && len(*stack) >= 0 {
// 		pop(stack)
// 		return
// 	}
// 	currentData := (*stack)[len(*stack)-1]
// 	pop(stack)
// 	delete(stack, val)
// 	push(stack, currentData)
// }

// func main() {
// 	var stack []int
// 	push(&stack, 8)
// 	push(&stack, 3)
// 	push(&stack, 9)
// 	push(&stack, 1)
// 	fmt.Println("@@", stack)
// 	// pop(&stack)
// 	delete(&stack, 3)
// 	delete(&stack, 8)
// 	fmt.Println("**", stack)
// }
