package BinaryHeap

type( 
	HeapType string
	BinaryHeap struct {
		data *[]int
		totalItems int
		heapType HeapType
	}
	BinaryHeapMethods interface {
		CreateBinaryHeap()
		ChildCheck(int, int) bool
		Heapify(int)
	}
)

func Swap(aIndex int, bIndex int, array *[]int) {
	temp := (*array)[aIndex]
	(*array)[aIndex] = (*array)[bIndex]
	(*array)[bIndex] = temp
}

func (heap *BinaryHeap) ChildCheck(index int, largestItemIndex int) bool{
	if index < heap.totalItems {
		if heap.heapType == "MAX" {
			return (*heap.data)[largestItemIndex] < (*heap.data)[index]
		} else if heap.heapType == "MIN" {
			return (*heap.data)[largestItemIndex] > (*heap.data)[index]
		}
	}
	return false
}

func (heap *BinaryHeap)Heapify(nodeIndex int) {
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

func (heap *BinaryHeap) CreateBinaryHeap() {
	for i:= heap.totalItems/2; i>=0; i-- {
		heap.Heapify(i)
	}
}
