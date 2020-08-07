package stack

import (
	"github.com/wednesday-solutions/golang-algorithm-club/pkg/utl"
	"testing"
)

var stack Stack

func createStack() {
	stack = NewStackFromArray([]string{"1", "2", "3", "4"})
}
func TestMain(m *testing.M) {
	utl.SetupTests(m, createStack)
}

func TestNewStackFromArray(t *testing.T) {
	stack := NewStackFromArray([]string{"1", "2"})
	if len(stack) != 2 {
		t.Errorf("length of stack should be 2 ")
	}
}

func TestNewStack(t *testing.T) {
	stack := NewStack()
	if len(stack) != 0 {
		t.Errorf("length of stack should be 0")
	}
}
func TestStack_Push(t *testing.T) {
	previousLength := len(stack)
	element := "1"
	stack.Push(element)
	lastElement := stack.Peek()
	if previousLength+1 != len(stack) {
		t.Errorf("Push should add an element to the stack thus increasing the stack size by 1")
	}
	if lastElement != element {
		t.Errorf("Push should add the new element to the top of the stack")
	}
}
func TestStack_Peek(t *testing.T) {
	element := "1"
	stack.Push(element)
	lastElement := stack.Peek()
	if lastElement != element {
		t.Errorf("Peek should return last added element")
	}
}

func TestStack_Pop(t *testing.T) {
	element := "1"
	stack.Push(element)
	previousLength := len(stack)
	lastElement := stack.Pop()
	if lastElement != element {
		t.Errorf("Pop should return last added element")
	}
	if previousLength != len(stack)+1 {
		t.Errorf("Pop should remove the last element from the stack")
	}
}
