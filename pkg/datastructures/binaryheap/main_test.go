package BinaryHeap

import (
	"github.com/wednesday-solutions/golang-algorithm-club/pkg/utl"
	"testing"
	"reflect"
)

func TestMain(m *testing.M) {
	utl.SetupTests(m, InitBinaryHeap)
}

var minData = []int {0, 1, 4, 2 ,3, 5, 6, 7, 8, 9}
var maxData = []int {0, 1, 4, 2 ,3, 5, 6, 7, 8, 9}
var maxHeap BinaryHeapMethods = &BinaryHeap {&maxData, 10, MAX}
var minHeap BinaryHeapMethods = &BinaryHeap {&minData, 10, MIN}

var minBinaryHeapResult *[]int
var maxBinaryHeapResult *[]int
var expectedMinBinaryHeapResult = []int {0, 1, 4, 2, 3, 5, 6, 7, 8, 9}
var expectedMaxBinaryHeapResult = []int {9, 8, 6, 7, 3, 5, 4, 0, 2, 1}

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