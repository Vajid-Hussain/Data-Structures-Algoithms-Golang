// package main

// import "fmt"

// type heapTree struct {
// 	heap []int
// 	size int
// }

// func (h *heapTree) insert(data int) {
// 	h.heap = append(h.heap, data)
// 	currentPos := len(h.heap) - 1
// 	parent := (currentPos - 1) / 2

// 	for currentPos >= 0 && h.heap[currentPos] > h.heap[parent] {
// 		h.heap[currentPos], h.heap[parent] = h.heap[parent], h.heap[currentPos]
// 		currentPos = parent
// 		parent = (currentPos - 1) / 2
// 	}
// }

// func (h *heapTree) preDelete() {
// 	h.size = len(h.heap) - 1
// 	for i := 1; i < len(h.heap); i++ {
// 		h.heap[len(h.heap)-1], h.heap[0] = h.heap[0], h.heap[len(h.heap)-1]
// 		h.size--
// 		h.delete(0)
// 	}
// }

// func (h *heapTree) delete(index int) {
// 	for {
// 		point := index
// 		leftNode := (point * 2) + 1
// 		rightNode := point*2 + 2

// 		if leftNode < h.size && h.heap[leftNode] > h.heap[point] {
// 			point = leftNode
// 		}
// 		if rightNode < h.size && h.heap[rightNode] > h.heap[point] {
// 			point = rightNode
// 		}
// 		if point == index {
// 			break
// 		}

// 		index = point
// 	}
// }

// func main() {
// 	heap := heapTree{}
// 	heap.insert(3)
// 	heap.insert(9)
// 	heap.insert(1)
// 	heap.insert(5)
// 	heap.insert(4)
// 	fmt.Println("**", heap.heap)
// 	heap.preDelete()
// 	fmt.Println(heap.heap)
// }
