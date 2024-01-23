// package main

// import "fmt"

// type MaxHeap struct {
// 	slice    []int
// 	heapSize int
// }

// func (h *MaxHeap) sort() {
// 	length := len(h.slice)

// 	// Build max heap
// 	for i := length/2; i >= 0; i-- {
// 		h.heapify(length, i)
// 		fmt.Println("@", h.slice)
// 	}

// 	// Extract elements one by one from the heap and swap with the root
// 	for i := length - 1; i >= 0; i-- {
// 		h.slice[0], h.slice[i] = h.slice[i], h.slice[0]
// 		h.heapify(i, 0)
// 	}
// }

// func (h *MaxHeap) heapify(length, i int) {
// 	largest := i
// 	left := 2*i + 1
// 	right := 2*i + 2

// 	// Check if left child is larger than the current largest
// 	if left < length && h.slice[left] > h.slice[largest] {
// 		largest = left
// 	}

// 	// Check if right child is larger than the current largest
// 	if right < length && h.slice[right] > h.slice[largest] {
// 		largest = right
// 	}

// 	// If the largest is not the current node, swap and recursively heapify
// 	if largest != i {
// 		h.slice[i], h.slice[largest] = h.slice[largest], h.slice[i]
// 		h.heapify(length, largest)
// 	}
// }

// func main() {
// 	maxHeap := &MaxHeap{slice: []int{2,5,9,6,1,4,0,3}}
// 	maxHeap.sort()
// 	fmt.Println(maxHeap.slice) // The sorted array

// }
