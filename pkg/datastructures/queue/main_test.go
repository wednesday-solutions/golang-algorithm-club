package queue

import (
	"github.com/wednesday-solutions/golang-algorithm-club/pkg/utl"
	"testing"
)

var queue Queue

func createQueue() {
	queue = NewQueueFromArray([]string{"1", "2", "3", "4"})
}

func TestMain(m *testing.M) {
	utl.SetupTests(m, createQueue)
}

func TestNewQueue(t *testing.T) {
	queue := NewQueue()
	if len(queue) > 0 {
		t.Errorf("length of new queue should be 0")
	}
}

func TestNewQueueFromArray(t *testing.T) {
	queue := NewQueueFromArray([]string{"1"})
	if len(queue) != 1 {
		t.Errorf("length of the new queue should be 1")
	}
}

func TestQueue_Dequeue(t *testing.T) {
	firstElement := queue.Peek()
	dequeuedElement := queue.Dequeue()
	if dequeuedElement != firstElement {
		t.Errorf("the first element should be dequeued")
	}
}

func TestQueue_DequeueEmpty(t *testing.T) {
	queue = NewQueue()
	peekedElement := queue.Dequeue()
	if peekedElement != "" {
		t.Errorf("for an empty queue Dequeue should return \"\"")
	}

}
func TestQueue_Enqueue(t *testing.T) {
	element := "5"
	previousLength := len(queue)
	enqueuedElement := queue.Enqueue(element)
	if enqueuedElement != element {
		t.Errorf("enqueued element should return the newly added element")
	}
	if previousLength+1 != len(queue) {
		t.Errorf("the length of the queue should increase by one")
	}
}

func TestQueue_Peek(t *testing.T) {
	firstElement := "1"
	secondElement := "2"
	queue := NewQueueFromArray([]string{firstElement, secondElement})
	previousLength := len(queue)
	peekedElement := queue.Peek()
	if peekedElement != firstElement {
		t.Errorf("Peek should return the first element")
	}
	if len(queue) != previousLength {
		t.Errorf("Peek should not remove the first element from the array")
	}
}

func TestQueue_PeekEmpty(t *testing.T) {
	queue = NewQueue()
	peekedElement := queue.Peek()
	if peekedElement != "" {
		t.Errorf("for an empty queue Peek should return \"\"")
	}

}
