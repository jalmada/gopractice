package utils

import (
	"fmt"
)

//Heap Classic Heap
type Heap struct {
	items []int
	tp    bool
}

//NewMinHeap Creates a new Heap
func NewMinHeap() *Heap {
	heap := Heap{}
	heap.tp = false
	//heap.items = make([]int, 0)
	heap.items = []int{10, 15, 20, 17}
	return &heap
}

//Peek the items array and retun the root value
func (heap *Heap) Peek() (int, bool) {
	if len(heap.items) == 0 {
		return -1, true
	}

	return heap.items[0], false
}

func (heap *Heap) swap(index1 int, index2 int) {
	tmp := heap.items[index1]
	heap.items[index1] = heap.items[index2]
	heap.items[index2] = tmp
}

//Poll the root and reorganizes the heap
func (heap *Heap) Poll() (int, bool) {
	if len(heap.items) == 0 {
		return -1, true
	}

	item := heap.items[0]
	heap.items[0] = heap.items[len(heap.items)-1]
	heap.items = heap.items[:len(heap.items)-1]
	heap.heapifyDown()
	return item, false
}

func (heap *Heap) heapifyDown() {
	index := 0
	for heap.hasLeftChild(index) {
		smallerChildIndex := heap.getLeftChildIndex(index)
		if heap.hasRightChild(index) && heap.leftChild(index) > heap.rightChild(index) {
			smallerChildIndex = heap.getRightChildIndex(index)
		}

		if heap.items[index] < heap.items[smallerChildIndex] {
			break
		} else {
			heap.swap(index, smallerChildIndex)
		}

		index = smallerChildIndex
	}

}

func (heap *Heap) heapifyUp() {
	index := len(heap.items) - 1

	for heap.hasParent(index) && heap.parent(index) > heap.items[index] {
		heap.swap(heap.getParentIndex(index), index)
		index = heap.getParentIndex(index)
	}
}

//Add a new item and reorganized
func (heap *Heap) Add(item int) {
	heap.items = append(heap.items, item)
	heap.heapifyUp()
}

func (heap *Heap) getLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (heap *Heap) getRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (heap *Heap) getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (heap *Heap) hasLeftChild(index int) bool {
	return heap.getLeftChildIndex(index) < len(heap.items)
}

func (heap *Heap) hasRightChild(index int) bool {
	return heap.getRightChildIndex(index) < len(heap.items)
}

func (heap *Heap) hasParent(index int) bool {
	return heap.getParentIndex(index) >= 0
}

func (heap *Heap) leftChild(index int) int {
	return heap.items[heap.getLeftChildIndex(index)]
}

func (heap *Heap) rightChild(index int) int {
	return heap.items[heap.getRightChildIndex(index)]
}

func (heap *Heap) parent(index int) int {
	return heap.items[heap.getParentIndex(index)]
}

//Print Heap
func (heap *Heap) Print() {
	fmt.Println(heap.items)
}
