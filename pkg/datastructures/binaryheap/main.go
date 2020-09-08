package binaryheap

type (
	// HeapType Set the binary heap type
	HeapType string
	// BinaryHeap provide the data array pointer, totalItems and the heap type.
	BinaryHeap struct {
		data       *[]int
		totalItems int
		heapType   HeapType
	}
	// BinaryHeaper Interface for BinaryHeap type.
	BinaryHeaper interface {
		CreateBinaryHeap()
		ChildCheck(int, int) bool
		Heapify(int)
	}
)

// MIN set the heap type as min
var MIN = HeapType("MIN")

// MAX set the heap type as max
var MAX = HeapType("MAX")

// Swap two items in the array
func Swap(aIndex int, bIndex int, array *[]int) {
	(*array)[aIndex], (*array)[bIndex] = (*array)[bIndex], (*array)[aIndex]
}

// ChildCheck check if a child node is larger or smaller than the item at the index provided
func (heap *BinaryHeap) ChildCheck(index int, largestItemIndex int) bool {
	if index < heap.totalItems {
		if heap.heapType == MAX {
			return (*heap.data)[largestItemIndex] < (*heap.data)[index]
		} else if heap.heapType == MIN {
			return (*heap.data)[largestItemIndex] > (*heap.data)[index]
		}
	}
	return false
}

// Heapify Swap the children and node until the heap condition is satisfied
func (heap *BinaryHeap) Heapify(nodeIndex int) {
	leftChildIndex := 2*nodeIndex + 1
	rightChildIndex := 2*nodeIndex + 2
	largest := nodeIndex
	if heap.ChildCheck(leftChildIndex, largest) {
		largest = leftChildIndex
	}
	if heap.ChildCheck(rightChildIndex, largest) {
		largest = rightChildIndex
	}

	if largest != nodeIndex {
		Swap(largest, nodeIndex, heap.data)
		heap.Heapify(largest)
	}
}

// CreateBinaryHeap Create a Binary heap
func (heap *BinaryHeap) CreateBinaryHeap() {
	for i := heap.totalItems / 2; i >= 0; i-- {
		heap.Heapify(i)
	}
}
