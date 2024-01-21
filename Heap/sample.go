package main

import "fmt"

type minHeap struct {
	heap []int
}

func (h *minHeap) insert(data int) {
	h.heap = append(h.heap, data)
	lastValue := len(h.heap) - 1
	for lastValue > 0 {
		parent := lastValue / 2
		if h.heap[parent] > h.heap[lastValue] {
			h.heap[parent], h.heap[lastValue] = h.heap[lastValue], h.heap[parent]
			lastValue = parent
		} else {
			return
		}
	}
}

func main() {
	heap := &minHeap{}
	heap.insert(4)
	heap.insert(5)
	heap.insert(2)
	heap.insert(8)
	heap.insert(1)
	fmt.Println("@@", heap.heap)
}
