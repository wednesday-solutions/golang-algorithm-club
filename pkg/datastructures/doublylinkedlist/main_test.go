package doublylinkedlist

import (
	"github.com/wednesday-solutions/golang-algorithm-club/pkg/utl"
	"testing"
)

var doublyLinkedList *DoublyLinkedList

func createDoublyLinkedList() {
	if doublyLinkedList != nil {
		doublyLinkedList.RemoveAll()
	}
	doublyLinkedList = CreateDoublyLinkedList("1")
	_, _ = doublyLinkedList.Append("2")
	_, _ = doublyLinkedList.Insert("3", 1)
}
func TestMain(m *testing.M) {
	utl.SetupTests(m, createDoublyLinkedList)
}

func TestCreateDoublyLinkedList(t *testing.T) {
	if doublyLinkedList.Head == nil {
		t.Errorf("doublyLinkedList.Head should not be nil")
	}
	if doublyLinkedList.Last == nil {
		t.Errorf("doublyLinkedList.Last should not be nil")
	}
	if doublyLinkedList.Length == 0 {
		t.Errorf("doublyLinkedList should have Length 1")
	}
	if doublyLinkedList.Node != nil {
		t.Errorf("doublyLinkedList.Node should be nil")
	}
}

func TestDoublyLinkedList_Append(t *testing.T) {
	newValue := "99"
	previousLength := doublyLinkedList.Length
	previousLast := doublyLinkedList.Last
	node, _ := doublyLinkedList.Append(newValue)
	if node == nil {
		t.Errorf("new node should be appended at the end of the Doubly Linked List")
	}
	if node != nil && node.Value != newValue {
		t.Errorf("new node should have the the value %s but got the value %s", newValue, node.Value)
	}
	if previousLength+1 != doublyLinkedList.Length {
		t.Errorf("length of the newly doubly Linked List should increase by one")
	}
	if doublyLinkedList.Last.Value != newValue {
		t.Errorf("value of the doublyLinkedList.Last should be the same as the value of the newly appended node")
	}
	if doublyLinkedList.Last.Previous != previousLast {
		t.Errorf("doublyLinkedList.Last.Previous should point to the newly appended node ")
	}
	if previousLast.Next != node {
		t.Errorf("previousLast.Next should point to the newly appended node ")
	}
}

func TestDoublyLinkedList_AppendEmptyHead(t *testing.T) {
	doublyLinkedList.RemoveAll()
	newValue := "99"
	node, _ := doublyLinkedList.Append(newValue)
	if doublyLinkedList.Head != node {
		t.Errorf("doublyLinkedList.Head should be equal to the newly added node")
	}
	if doublyLinkedList.Last != node {
		t.Errorf("doublyLinkedList.Head should be equal to the newly added node")
	}
	createDoublyLinkedList()
}

func TestDoublyLinkedList_Count(t *testing.T) {
	doublyLinkedListArr, _ := doublyLinkedList.ToArray()
	if len(doublyLinkedListArr) != doublyLinkedList.Count() {
		t.Errorf("doublyLinkedList.Count() should be equal to the total number of nodes")
	}
	if doublyLinkedList.Length != doublyLinkedList.Count() {
		t.Errorf("doublyLinkedList.Count() should be equal to the total number of nodes")
	}
}

func TestDoublyLinkedList_Insert(t *testing.T) {
	newValue := "44"
	insertAtIndex := 1
	nodeAt1, _ := doublyLinkedList.NodeAt(insertAtIndex)
	previousNode := nodeAt1.Previous

	node, _ := doublyLinkedList.Insert(newValue, insertAtIndex)
	if node == nil {
		t.Errorf("doublyLinkedList.Insert() should create new node")
	}

	if node != previousNode.Next {
		t.Errorf("node should should be equal to previousNode.Next")
	}
	if node != nil && node.Previous != previousNode {
		t.Errorf("nodeAt1.Previous should should be equal to previousNode")
	}
	if node != previousNode.Next {
		t.Errorf("node should should be equal to previousNode.Next")
	}
	if node.Next != nodeAt1 {
		t.Errorf("node.Next should should be equal to nodeAt1")
	}
	if nodeAt1.Previous != node {
		t.Errorf("nodeAt1.Previous should should be equal to the newly added node")
	}
}
func TestDoublyLinkedList_InsertToEmpty(t *testing.T) {
	doublyLinkedList.RemoveAll()
	node, _ := doublyLinkedList.Insert("0", 0)
	if doublyLinkedList.Head != node {
		t.Errorf("doublyLinkedList.Head should be equal to the newly inserted node")
	}
	if doublyLinkedList.Last != node {
		t.Errorf("doublyLinkedList.Last should be equal to the newly inserted node")
	}
	createDoublyLinkedList()
}
func TestDoublyLinkedList_InsertAtZero(t *testing.T) {
	previousHead := doublyLinkedList.Head
	node, _ := doublyLinkedList.Insert("0", 0)

	if doublyLinkedList.Head != node {
		t.Errorf("doublyLinkedList.Head should be equal to the newly inserted node")
	}
	if node == previousHead.Next {
		t.Errorf("previousHead.Next should be equal to the newly inserted node")
	}
	if node.Previous == previousHead {
		t.Errorf("previousHead should be equal to the previousHead")
	}
}

func TestDoublyLinkedList_InsertAtInvalidIndex(t *testing.T) {
	node, _ := doublyLinkedList.Insert("0", 100)
	if node != nil {
		t.Errorf("Insert should fail due to invalid index")
	}
}

func TestDoublyLinkedList_NodeAt(t *testing.T) {
	nodeToFind := doublyLinkedList.Head.Next
	node, _ := doublyLinkedList.NodeAt(1)
	if node != nodeToFind {
		t.Errorf("node should be equal to doublyLinkedList.Head.Next")
	}
}

func TestDoublyLinkedList_NodeAtInvalidIndex(t *testing.T) {
	node, err := doublyLinkedList.NodeAt(999)
	if node != nil || err == nil {
		t.Errorf("should not be able to find a node since the index is invalid")
	}
}

func TestDoublyLinkedList_Remove(t *testing.T) {
	currentHead := doublyLinkedList.Head
	nextNode, _ := doublyLinkedList.NodeAt(2)

	_, _ = doublyLinkedList.Remove(1)
	if nextNode != currentHead.Next {
		t.Errorf("nextNode should be equal to currentHead.Next")
	}
	if nextNode.Previous != currentHead {
		t.Errorf("nextNode.Previous should be equal to currentHead")
	}
}

func TestDoublyLinkedList_RemoveInvalidIndex(t *testing.T) {
	_, err := doublyLinkedList.Remove(999)
	if err == nil {
		t.Errorf("should not be able to remove the node since the index is invalid")
	}
}

func TestDoublyLinkedList_RemoveAll(t *testing.T) {
	doublyLinkedList.RemoveAll()
	if doublyLinkedList.Head != nil {
		t.Errorf("doublyLinkedList.Head should be nil")
	}
	if doublyLinkedList.Last != nil {
		t.Errorf("doublyLinkedList.Last should be nil")
	}
	if doublyLinkedList.Length != 0 {
		t.Errorf("doublyLinkedList.Length should be 0")
	}
	createDoublyLinkedList()
}

func TestDoublyLinkedList_RemoveLast(t *testing.T) {
	createDoublyLinkedList()
	previousLast := doublyLinkedList.Last
	previousLength := doublyLinkedList.Length
	_, _ = doublyLinkedList.RemoveLast()
	if doublyLinkedList.Last == previousLast {
		t.Errorf("doublyLinkedList.Last should not be equal to previousLast")
	}
	if previousLength == doublyLinkedList.Length {
		t.Errorf("doublyLinkedList.Length should not be less than previousLength")
	}
}

func TestDoublyLinkedList_Reverse(t *testing.T) {
	doublyLinkedListArr, _ := doublyLinkedList.ToArray()
	doublyLinkedList = doublyLinkedList.Reverse()
	doublyLinkedListReversedArr, _ := doublyLinkedList.ToArray()
	for i := 0; i < doublyLinkedList.Length; i++ {
		if doublyLinkedListArr[i] != doublyLinkedListReversedArr[doublyLinkedList.Length-1-i] {
			t.Errorf("the doubly linked list has incorrect values at index: %d after reversal", i)
		}
	}

}

func TestDoublyLinkedList_ToArray(t *testing.T) {
	doublyLinkedListArr, _ := doublyLinkedList.ToArray()
	if len(doublyLinkedListArr) != doublyLinkedList.Count() {
		t.Errorf("length of array is not correct")
	}
	if len(doublyLinkedListArr) != doublyLinkedList.Length {
		t.Errorf("length of array is not correct")
	}
	element := doublyLinkedList.Head
	for i := 0; i < doublyLinkedList.Length; i++ {
		if element.Value != doublyLinkedListArr[i] {
			t.Errorf("incorrect value while iterating over doublyLinkedListArr at index: %d", i)
		}
		element = element.Next
	}
}

func TestDoublyLinkedList_Next(t *testing.T) {
	doublyLinkedListArr, _ := doublyLinkedList.ToArray()
	var doublyLinkedListIterableArr []string
	for {
		v, ok := doublyLinkedList.Next()
		if !ok {
			break
		}
		doublyLinkedListIterableArr = append(doublyLinkedListIterableArr, v.Value)
	}
	for i := 0; i < doublyLinkedList.Length; i++ {
		if doublyLinkedListIterableArr[i] != doublyLinkedListArr[i] {
			t.Errorf("incorrect value while iterating over doublyLinkedListArr at index: %d", i)
		}
	}

}
