// package main

// import "fmt"

// type heapTree struct {
// 	heap []int
// 	size int
// }

// func (h *heapTree) insert(data int) {
// 	h.heap = append(h.heap, data)
// 	pos := len(h.heap) - 1
// 	parent := (pos - 1) / 2

// 	for pos >= 0 && h.heap[parent] < h.heap[pos] {
// 		h.heap[parent], h.heap[pos] = h.heap[pos], h.heap[parent]
// 		pos = parent
// 		parent = (pos - 1) / 2
// 	}
// }

// func (h *heapTree) preDelete() {
// 	h.heap[len(h.heap)-1], h.heap[0] = h.heap[0], h.heap[len(h.heap)-1]
// 	h.heap = h.heap[:len(h.heap)-1]
// 	h.size = len(h.heap) - 1
// 	h.delete(0)
// }

// func (h *heapTree) sort() {
// 	for i := len(h.heap) / 2; i >= 0; i-- {
// 		h.delete(i)
// 	}
// 	for i := 0; i < len(h.heap); i++ {
// 		h.heap[h.size], h.heap[0] = h.heap[0], h.heap[h.size]
// 		h.size--
// 		h.delete(0)
// 	}
// }

// func (h *heapTree) delete(index int) {
// 	largest := index
// 	left := index*2 + 1
// 	right := index*2 + 2
// 	if left <= h.size && h.heap[left] > h.heap[largest] {
// 		largest = left
// 	}
// 	if right <= h.size && h.heap[right] > h.heap[largest] {
// 		largest = right
// 	}
// 	if largest != index {
// 		h.heap[largest], h.heap[index] = h.heap[index], h.heap[largest]
// 		h.delete(largest)
// 	}
// }

// func main() {
// 	heap := heapTree{}
// 	heap.insert(2)
// 	heap.insert(8)
// 	heap.insert(5)
// 	heap.insert(1)
// 	heap.insert(7)
// 	fmt.Println(heap.heap)
// 	heap.preDelete()
// 	fmt.Println(heap.heap)
// 	// heap.heap=[]int{4,2,7,9,1}
// 	heap.sort()
// 	fmt.Println(heap.heap)
// }
