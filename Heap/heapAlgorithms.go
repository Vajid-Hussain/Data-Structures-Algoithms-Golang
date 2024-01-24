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
	hp.size = len(hp.heap) - 1
	for i := len(hp.heap) / 2; i >= 0; i-- {
		hp.maxheapify(i)
	}
	fmt.Println("max heap", hp.heap)
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
	// heap.sortAccesnding()

	// fmt.Println(heap.heap)
}
