package linkedlist

import (
	"github.com/wednesday-solutions/golang-algorithm-club/pkg/utl"
	"testing"
)

var linkedList *LinkedList

func createLinkedList() {
	if linkedList != nil {
		linkedList.RemoveAll()
	}
	linkedList = CreateLinkedList("1")
	_, _ = linkedList.Append("2")
	_, _ = linkedList.Insert("3", 1)
}
func TestMain(m *testing.M) {
	utl.SetupTests(m, createLinkedList)
}

func TestCreateLinkedList(t *testing.T) {
	linkedList := CreateLinkedList("1")
	if linkedList.Length != 1 {
		t.Errorf("created linked list should have length 1")
	}

	if linkedList.Head == nil {
		t.Errorf("head of linked list should not be empty")
	}
}

func TestLinkedList_Append(t *testing.T) {
	element := "4"
	previousLength := linkedList.Length
	addedElement, _ := linkedList.Append(element)
	if addedElement.Value != element {
		t.Errorf("value of added element should be equal to the new element")
	}
	if linkedList.Length != previousLength+1 {
		t.Errorf("length of the linked list should increase by one")
	}
}

func TestLinkedList_AppendEmpty(t *testing.T) {
	linkedList.RemoveAll()
	element := "4"
	addedElement, _ := linkedList.Append(element)
	if linkedList.Head != addedElement {
		t.Errorf("for an empty linked list newly appended element should be the head ")
	}
}

func TestLinkedList_Count(t *testing.T) {
	createLinkedList()
	if linkedList.Count() != 3 {
		t.Errorf("Count should return the total number of items in the linked list")
	}

	if linkedList.Count() != linkedList.Length {
		t.Errorf("count should be equal to the lin")
	}
}

func TestLinkedList_GetLast(t *testing.T) {
	firstElement := "1"
	secondElement := "2"
	linkedList := CreateLinkedListFromArray([]string{firstElement, secondElement})
	if linkedList.GetLast().Value != secondElement {
		t.Errorf("GetLast should return the last element")
	}
}

func TestCreateLinkedListFromArray(t *testing.T) {
	firstElement := "1"
	secondElement := "2"
	linkedList := CreateLinkedListFromArray([]string{firstElement, secondElement})
	if linkedList.Head.Value != firstElement {
		t.Errorf("linkedList.Head should have the same value as the first element")
	}
	if linkedList.GetLast().Value != secondElement {
		t.Errorf("linkedList.GetLast should have the same value as the second element")
	}
	if linkedList.Count() != 2 {
		t.Errorf("linked list should have 2 element")
	}
}

func TestLinkedList_Insert(t *testing.T) {
	element := "99"
	previousLength := linkedList.Count()
	insertedElement, _ := linkedList.Insert(element, 1)
	if insertedElement.Value != element {
		t.Errorf("insert should return the newly created element")
	}
	if linkedList.Count() != previousLength+1 {
		t.Errorf("inserting a new element should increase the count by 1")
	}
	if linkedList.Head.Next != insertedElement {
		t.Errorf("insert at 1st index should add element as next of head")
	}
}

func TestLinkedList_InsertInvalidIndex(t *testing.T) {
	element := "99"
	insertedElement, err := linkedList.Insert(element, 999)
	if err == nil {
		t.Errorf("error should fail with invalid index error")
	}
	if insertedElement != nil {
		t.Errorf("insert should return nil as the first parameter")
	}
}

func TestLinkedList_InsertHead(t *testing.T) {
	element := "99"
	insertedElement, _ := linkedList.Insert(element, 0)
	if insertedElement != linkedList.Head {
		t.Errorf("inserting at the 0th index should add the new element at the head")
	}
	linkedList.RemoveAll()
	addedElement, _ := linkedList.Insert(element, 0)
	if addedElement != linkedList.Head {
		t.Errorf("insert on an empty linkedlist should add the new element at the head")
	}
	createLinkedList()
}

func TestLinkedList_NodeAt(t *testing.T) {
	nodeToFind := linkedList.Head.Next
	node, _ := linkedList.NodeAt(1)
	if node != nodeToFind {
		t.Errorf("node should be equal to linkedList.Head.Next")
	}
}

func TestLinkedList_NodeAtInvalidIndex(t *testing.T) {
	node, err := linkedList.NodeAt(999)
	if node != nil || err == nil {
		t.Errorf("should not be able to find a node since the index is invalid")
	}
}

func TestLinkedList_Remove(t *testing.T) {
	currentHead := linkedList.Head
	nextNode, _ := linkedList.NodeAt(2)

	_, _ = linkedList.Remove(1)
	if nextNode != currentHead.Next {
		t.Errorf("nextNode should be equal to currentHead.Next")
	}

}

func TestLinkedList_RemoveInvalidIndex(t *testing.T) {
	_, err := linkedList.Remove(999)
	if err == nil {
		t.Errorf("should not be able to remove the node since the index is invalid")
	}
}

func TestLinkedList_RemoveAll(t *testing.T) {
	linkedList.RemoveAll()
	if linkedList.Head != nil {
		t.Errorf("linkedList.Head should be nil")
	}
	if linkedList.GetLast() != nil {
		t.Errorf("linkedList.GetLast() should be nil")
	}
	if linkedList.Length != 0 {
		t.Errorf("linkedList.Length should be 0")
	}
	createLinkedList()
}

func TestLinkedList_RemoveLast(t *testing.T) {
	createLinkedList()
	previousLast := linkedList.GetLast()
	previousLength := linkedList.Length
	_, _ = linkedList.RemoveLast()
	if linkedList.GetLast() == previousLast {
		t.Errorf("linkedList.Last should not be equal to previousLast")
	}
	if previousLength == linkedList.Length {
		t.Errorf("linkedList.Length should not be less than previousLength")
	}
}

func TestLinkedList_Reverse(t *testing.T) {
	linkedListListArr, _ := linkedList.ToArray()
	linkedList = linkedList.Reverse()
	linkedListListReversedArr, _ := linkedList.ToArray()
	for i := 0; i < linkedList.Length; i++ {
		if linkedListListArr[i] != linkedListListReversedArr[linkedList.Length-1-i] {
			t.Errorf("the doubly linked list has incorrect values at index: %d after reversal", i)
		}
	}

}

func TestLinkedList_ToArray(t *testing.T) {
	linkedListListArr, _ := linkedList.ToArray()
	if len(linkedListListArr) != linkedList.Count() {
		t.Errorf("length of array is not correct")
	}
	if len(linkedListListArr) != linkedList.Length {
		t.Errorf("length of array is not correct")
	}
	element := linkedList.Head
	for i := 0; i < linkedList.Length; i++ {
		if element.Value != linkedListListArr[i] {
			t.Errorf("incorrect value while iterating over linkedListListArr at index: %d", i)
		}
		element = element.Next
	}
}

func TestLinkedList_Next(t *testing.T) {
	linkedListListArr, _ := linkedList.ToArray()
	var linkedListListIterableArr []string
	for {
		v, ok := linkedList.Next()
		if !ok {
			break
		}
		linkedListListIterableArr = append(linkedListListIterableArr, v.Value)
	}
	for i := 0; i < linkedList.Length; i++ {
		if linkedListListIterableArr[i] != linkedListListArr[i] {
			t.Errorf("incorrect value while iterating over linkedListListArr at index: %d", i)
		}
	}

}
