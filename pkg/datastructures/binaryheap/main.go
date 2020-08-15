package BinaryHeap

type HeapConditions struct {
	heapType string
	totalItems int
}

var CreateHeap HeapConditions

func swap(aIndex int, bIndex int, array *[]int) {
	temp := (*array)[aIndex]
	(*array)[aIndex] = (*array)[bIndex]
	(*array)[bIndex] = temp
}

func childCheck(index int, largestItemIndex int, array * []int) bool{
	if index < CreateHeap.totalItems {
		if CreateHeap.heapType == "MAX" {
			return (*array)[largestItemIndex] < (*array)[index]
		} else if CreateHeap.heapType == "MIN" {
			return (*array)[largestItemIndex] > (*array)[index]
		}
	}
	return false
}

func heapify(array *[]int, nodeIndex int) {
	leftChildIndex := 2*nodeIndex + 1
	rightChildIndex := 2*nodeIndex + 2
	largest := nodeIndex

	if childCheck(leftChildIndex, largest, array) {
		largest = leftChildIndex
	}
	if childCheck(rightChildIndex, largest, array) {
		largest = rightChildIndex
	}

	if largest != nodeIndex {
		swap (largest, nodeIndex, array)
		heapify(array, largest)
	}
}

func createBinaryHeap(array []int, n int, heapType string) *[]int {
	CreateHeap = HeapConditions {heapType, n}
	for i:= n/2; i>=0; i-- {
		heapify(&array, i)
	}
	return &array
}
