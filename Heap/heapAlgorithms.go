package main

import "fmt"

type heapTree struct {
	heap []int
	size int
}

func (hp *heapTree) insert(data int) {
	hp.heap = append(hp.heap, data)
	index := len(hp.heap) - 1
	parent := (index - 1) / 2

	for index >= 0 && hp.heap[parent] < hp.heap[index] {
		hp.heap[parent], hp.heap[index] = hp.heap[index], hp.heap[parent]
		index = parent
		parent = (index - 1) / 2
	}
}

func (hp *heapTree) delete() {
	hp.size = len(hp.heap) - 1
	hp.heap[len(hp.heap)-1], hp.heap[0] = hp.heap[0], hp.heap[len(hp.heap)-1]
	hp.heap = hp.heap[:len(hp.heap)-1]
	hp.size--
	hp.maxheapify(0)
}

func (hp *heapTree) sortAccesnding() {
	// fmt.Println("===", hp.heap)
	hp.size = len(hp.heap) - 1
	for i := len(hp.heap) / 2; i >= 0; i-- {
		hp.maxheapify(i)
	}
	// fmt.Println("max heap", hp.heap)
	// hp.size = len(hp.heap) - 1
	for i := hp.size; i >= 0; i-- {
		hp.heap[i], hp.heap[0] = hp.heap[0], hp.heap[i]
		hp.size--
		hp.maxheapify(0)
	}
}

func (hp *heapTree) maxheapify(index int) {
	largest := index
	left := largest*2 + 1
	right := largest*2 + 2
	if left <= hp.size && hp.heap[left] > hp.heap[largest] {
		largest = left
	}
	if right <= hp.size && hp.heap[right] > hp.heap[largest] {
		largest = right
	}
	if largest != index {
		hp.heap[largest], hp.heap[index] = hp.heap[index], hp.heap[largest]
		hp.maxheapify(largest)
	}
}

func (hp *heapTree) minHeapInsert(data int) {
	hp.heap = append(hp.heap, data)
	index := len(hp.heap) - 1
	parent := (index - 1) / 2
	for index >= 0 && hp.heap[parent] > hp.heap[index] {
		hp.heap[parent], hp.heap[index] = hp.heap[index], hp.heap[parent]
		index = parent
		parent = (index - 1) / 2
	}
}

func (hp *heapTree) minheapDelete() {
	hp.size = len(hp.heap) - 1
	hp.heap[hp.size], hp.heap[0] = hp.heap[0], hp.heap[hp.size]
	hp.heap = hp.heap[:hp.size]
	hp.size--
	hp.minheapyfy(0)
}

func (hp *heapTree) sortDecending() {
	fmt.Println("@@", hp.heap)
	hp.size = len(hp.heap) - 1
	for i := hp.size / 2; i >= 0; i-- {
		hp.minheapyfy(i)
	}
	fmt.Println("**", hp.heap)
	for i := hp.size; i >= 0; i-- {
		hp.heap[i], hp.heap[0] = hp.heap[0], hp.heap[i]
		hp.size--
		hp.minheapyfy(0)
	}
}

func (hp *heapTree) minheapyfy(index int) {
	largest := index
	left := index*2 + 1
	right := index*2 + 2

	if hp.size >= left && hp.heap[largest] > hp.heap[left] {
		largest = left
	}
	if hp.size >= right && hp.heap[largest] > hp.heap[right] {
		largest = right
	}
	if index != largest {
		hp.heap[largest], hp.heap[index] = hp.heap[index], hp.heap[largest]
		hp.minheapyfy(largest)
	}
}

func main() {
	heap := heapTree{}
	heap.insert(4)
	heap.insert(2)
	heap.insert(6)
	heap.insert(5)
	heap.insert(1)
	heap.insert(29)
	heap.insert(7)
	fmt.Println(heap.heap)
	// heap.delete()
	// heap.sortAccesnding()
	heap.sortDecending()
	fmt.Println("end result after deletion:", heap.heap)

	// heap.minHeapInsert(5)
	// heap.minHeapInsert(2)
	// heap.minHeapInsert(9)
	// heap.minHeapInsert(7)
	// heap.minHeapInsert(6)
	// heap.minHeapInsert(3)
	// fmt.Println("minheap:", heap.heap)
	// heap.minheapDelete()
	heap.sortAccesnding()

	fmt.Println(heap.heap)
}


// package main

// import "fmt"

// type heapTree struct {
// 	heap []int
// 	size int
// }

// func (h *heapTree) MaxHeap(data int) {
// 	h.heap = append(h.heap, data)
// 	h.size = len(h.heap) - 1
// 	child := len(h.heap) - 1
// 	parent := (child - 1) / 2

// 	for parent >= 0 && h.heap[parent] < h.heap[child] {
// 		h.heap[parent], h.heap[child] = h.heap[child], h.heap[parent]
// 		child = parent
// 		parent = (parent - 1) / 2
// 	}
// }

// func (h *heapTree) HeapifyDownMax(index int) {
// 	largest := index
// 	left := index*2 + 1
// 	right := index*2 + 2
// 	if h.size >= left && h.heap[left] > h.heap[largest] {
// 		largest = left
// 	}
// 	if h.size >= right && h.heap[right] > h.heap[largest] {
// 		largest = right
// 	}

// 	if largest != index {
// 		h.heap[largest], h.heap[index] = h.heap[index], h.heap[largest]
// 		h.HeapifyDownMax(largest)
// 	}
// }

// func (h *heapTree) sortAccending() {
// 	for i := len(h.heap) / 2; i >= 0; i-- {
// 		h.HeapifyDownMax(i)
// 	}

// 	h.size = len(h.heap) - 1
// 	for i := len(h.heap) - 1; i >= 0; i-- {
// 		h.heap[0], h.heap[i] = h.heap[i], h.heap[0]
// 		h.size--
// 		h.HeapifyDownMax(0)
// 	}
// }

// func (h *heapTree) HeapifyDownMin(index int) {
// 	smallest := index
// 	left := index*2 + 1
// 	right := index*2 + 2
// 	if h.size >= left && h.heap[left] < h.heap[smallest] {
// 		smallest = left
// 	}
// 	if h.size >= right && h.heap[right] < h.heap[smallest] {
// 		smallest = right
// 	}
// 	if index != smallest {
// 		h.heap[smallest], h.heap[index] = h.heap[index], h.heap[smallest]
// 		h.HeapifyDownMin(smallest)
// 	}
// }

// func (h *heapTree) sortDesending() {
// 	for i := len(h.heap) / 2; i >= 0; i-- {
// 		h.HeapifyDownMin(i)
// 	}
// 	h.size = len(h.heap) - 1

// 	for i := len(h.heap) - 1; i >= 0; i-- {
// 		h.heap[0], h.heap[i] = h.heap[i], h.heap[0]
// 		h.size--
// 		h.HeapifyDownMin(0)
// 	}
// }

// func main() {
// 	heap := heapTree{}
// 	heap.MaxHeap(5)
// 	heap.MaxHeap(1)
// 	heap.MaxHeap(9)
// 	heap.MaxHeap(7)
// 	heap.MaxHeap(4)
// 	fmt.Println(heap.heap)
// 	heap.sortAccending()
// 	fmt.Println(heap.heap)

// 	heap.sortDesending()
// 	fmt.Println(heap.heap)
// }
