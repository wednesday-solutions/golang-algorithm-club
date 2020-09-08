package binaryheap

import (
	"github.com/wednesday-solutions/golang-algorithm-club/pkg/utl"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	utl.SetupTests(m, InitBinaryHeap)
}

var minData = []int{0, 1, 4, 2, 3, 5, 6, 7, 8, 9}
var maxData = []int{0, 1, 4, 2, 3, 5, 6, 7, 8, 9}
var maxHeap BinaryHeaper = &BinaryHeap{&maxData, 10, MAX}
var minHeap BinaryHeaper = &BinaryHeap{&minData, 10, MIN}
var expectedMinBinaryHeapResult = []int{0, 1, 4, 2, 3, 5, 6, 7, 8, 9}
var expectedMaxBinaryHeapResult = []int{9, 8, 6, 7, 3, 5, 4, 0, 2, 1}

func InitBinaryHeap() {
	maxHeap.CreateBinaryHeap()
	minHeap.CreateBinaryHeap()
}

func TestCreateMaxBinaryHeap(t *testing.T) {
	if !reflect.DeepEqual(maxData, expectedMaxBinaryHeapResult) {
		t.Errorf("The MaxBinaryHeap should be %v", expectedMaxBinaryHeapResult)
	}
}

func TestCreateMinBinaryHeap(t *testing.T) {
	if !reflect.DeepEqual(minData, expectedMinBinaryHeapResult) {
		t.Errorf("The MaxBinaryHeap should be %v", expectedMinBinaryHeapResult)
	}
}

func TestSwap(t *testing.T) {
	data := []int{1, 2}
	expectedData := []int{2, 1}
	Swap(0, 1, &data)
	if !reflect.DeepEqual(data, expectedData) {
		t.Errorf("data after swapping 0th and 1st index should be %v", expectedData)
	}
}

func TestChildCheck(t *testing.T) {
	maxZeroIndexCheck := maxHeap.ChildCheck(0, 1)
	if !maxZeroIndexCheck {
		t.Errorf("Index zero should have the greatest value in a MAX type heap")
	}
	maxNodeCheck := maxHeap.ChildCheck(1, 3)
	if !maxNodeCheck {
		t.Errorf("Child of a node should have a lower value in a MAX type heap")
	}
	minZeroIndexCheck := minHeap.ChildCheck(0, 1)
	if !minZeroIndexCheck {
		t.Errorf("Index zero should have the smallest value in a MIN type heap")
	}
	minNodeCheck := minHeap.ChildCheck(1, 3)
	if !minNodeCheck {
		t.Errorf("Child of a node should have a greater value in a MIN type heap")
	}
}

func TestHeapify(t *testing.T) {
	maxHeapifyCheck := []int{1, 2, 3}
	var newMaxHeap BinaryHeaper = &BinaryHeap{&maxHeapifyCheck, 3, MAX}
	newMaxHeap.Heapify(0)
	expectedMAXNode := []int{3, 2, 1}
	if !reflect.DeepEqual(expectedMAXNode, maxHeapifyCheck) {
		t.Errorf("Heapify should swap elements to satisfy MAX heap type")
	}
	minHeapifyCheck := []int{2, 1, 3}
	var newMinHeap BinaryHeaper = &BinaryHeap{&minHeapifyCheck, 3, MIN}
	newMinHeap.Heapify(0)
	expectedMINNode := []int{1, 2, 3}
	if !reflect.DeepEqual(expectedMINNode, minHeapifyCheck) {
		t.Errorf("Heapify should swap elements to satisfy MIN heap type")
	}
}
