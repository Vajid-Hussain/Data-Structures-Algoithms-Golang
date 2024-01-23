package main

import "fmt"

type heapTree struct {
	heap []int
	size int
}

func (h *heapTree) sort() {
	h.size = len(h.heap) - 1
	for i := len(h.heap) / 2; i >= 0; i-- {
		h.heapify(i)
	}
	fmt.Println(h.heap)

	for i:=0; i<len(h.heap); i++{
		h.heap[h.size], h.heap[0]= h.heap[0], h.heap[h.size]
		h.size--
		h.heapify(0)
	}
}

func (h *heapTree) heapify(index int) {
	largest := index
	left := largest*2 + 1
	right := largest*2 + 2

	if left <= h.size && h.heap[left] < h.heap[largest] {
		largest = left
	}
	if right <= h.size && h.heap[right] < h.heap[largest] {
		largest = right
	}
	if largest != index {
		h.heap[largest], h.heap[index] = h.heap[index], h.heap[largest]
		h.heapify(largest)
	}
}

func main() {
	hp := heapTree{heap: []int{3, 9, 6, 5, 1, 7, 30}}
	hp.sort()
	fmt.Println(hp.heap)
}
