package heap

//Heap Classic Heap
type Heap struct {
	capacity int
	size     int
	items    []int
	//Min = false, Max = true
	tp bool
}

//NewMinHeap Creates a new Heap
func NewMinHeap() *Heap {
	heap := Heap{}
	heap.capacity = 0
	heap.size = 0
	heap.tp = false
	heap.items = make([]int, heap.size, heap.capacity)

	return &heap
}

func (heap *Heap) peek() (int, bool) {
	if heap.size == 0 {
		return -1, true
	}

	return heap.items[0], false
}

func (heap *Heap) swap(index1 int, index2 int) {
	tmp := heap.items[index1]
	heap.items[index1] = heap.items[index2]
	heap.items[index2] = tmp
}

func (heap *Heap) ensureExtraCapacity() {
	if heap.size == heap.capacity {
		tmp := make([]int, len(heap.items), (cap(heap.items)+1)*2)
		for i := range heap.items {
			tmp[i] = heap.items[i]
		}
		heap.items = tmp
	}
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
	return heap.getLeftChildIndex(index) < heap.size
}

func (heap *Heap) hasRightChild(index int) bool {
	return heap.getRightChildIndex(index) < heap.size
}

func (heap *Heap) hasParent(index int) bool {
	return heap.getParentIndex(index) > 0
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
