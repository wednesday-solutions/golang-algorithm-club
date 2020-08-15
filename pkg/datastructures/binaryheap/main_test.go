package BinaryHeap

import (
	"github.com/wednesday-solutions/golang-algorithm-club/pkg/utl"
	"testing"
	"reflect"
)

func TestMain(m *testing.M) {
	utl.SetupTests(m, initCreateBinaryHeap)
}

var minBinaryHeapResult *[]int
var maxBinaryHeapResult *[]int
var expectedMinBinaryHeapResult = []int {0, 1, 4, 2, 3, 5, 6, 7, 8, 9}
var expectedMaxBinaryHeapResult = []int {9, 8, 6, 7, 3, 5, 4, 0, 2, 1}

func initCreateBinaryHeap() {
	array := []int {0, 1, 4, 2 ,3, 5, 6, 7, 8, 9}
	newArray := []int {0, 1, 4, 2 ,3, 5, 6, 7, 8, 9}
	minBinaryHeapResult = createBinaryHeap(array, 10, "MIN")
	maxBinaryHeapResult = createBinaryHeap(newArray, 10, "MAX")
}

func TestCreateMaxBinaryHeap(t *testing.T) {
	if !reflect.DeepEqual(*minBinaryHeapResult, expectedMinBinaryHeapResult) {
		t.Errorf("The MinBinaryHeap should be %v",expectedMinBinaryHeapResult)
	}
}

func TestCreateMinBinaryHeap(t *testing.T) {
	if !reflect.DeepEqual(*maxBinaryHeapResult, expectedMaxBinaryHeapResult) {
		t.Errorf("The MaxBinaryHeap should be %v", expectedMaxBinaryHeapResult)
	}
}