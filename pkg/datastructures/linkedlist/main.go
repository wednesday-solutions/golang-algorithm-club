package linkedlist

import (
	"fmt"
)

type (
	//Node Linked List node
	Node struct {
		Value string
		Next  *Node
	}
	//LinkedList Linked List struct
	LinkedList struct {
		LLInterface
		Head      *Node
		Node      *Node
		Length    int
		iterating bool
	}
	//LLInterface Interface for Linked List
	LLInterface interface {
		Append(value string) (*Node, error)
		Count() int
		Insert(value string, index int) (*Node, error)
		Next() (*Node, bool)
		NodeAt(index int) (*Node, error)
		Remove(index int) (string, error)
		RemoveAll()
		RemoveLast() (string, error)
		Reverse() *LinkedList
		ToArray() ([]string, error)
		GetLast() *Node
	}
)

func createLinkedListNode(value string) *Node {
	return &Node{
		Value: value,
		Next:  nil,
	}
}

func (linkedList *LinkedList) appendNode(node *Node) (*Node, error) {
	if linkedList.GetLast() == nil {
		linkedList.Head = node
	} else {
		linkedList.GetLast().Next = node
	}
	linkedList.Length++
	return node, nil
}
func (linkedList *LinkedList) insertNode(node *Node, index int) (*Node, error) {
	if linkedList.Head == nil {
		linkedList.Head = node
	} else if index < 0 || index > linkedList.Length {
		return nil, fmt.Errorf("invalid index")
	} else if index == 0 {
		node.Next = linkedList.Head
		linkedList.Head = node
	} else {
		nodeToReplace, err := linkedList.NodeAt(index)
		if err != nil {
			return nil, err
		}
		previousNode, err := linkedList.NodeAt(index - 1)
		if err == nil {
			previousNode.Next = node
		}
		node.Next = nodeToReplace
	}
	linkedList.Length++
	return node, nil
}

//GetLast get the last element in the linkedList
func (linkedList *LinkedList) GetLast() *Node {
	element := linkedList.Head
	for element != nil && element.Next != nil {
		element = element.Next
	}
	return element
}
func (linkedList *LinkedList) removeNode(node *Node) (*Node, error) {

	if linkedList.Length == 0 {
		return nil, fmt.Errorf("invalid node")
	}
	var element = linkedList.Head
	var previousElement = linkedList.Head
	for element != node {
		previousElement = element
		element = element.Next
	}
	previousElement.Next = node.Next
	return node, nil
}

//CreateLinkedList Create a new Linked List
func CreateLinkedList(value string) *LinkedList {
	node := createLinkedListNode(value)
	return &LinkedList{
		Head:   node,
		Length: 1,
		Node:   nil,
	}
}

//CreateLinkedListFromArray Create a new Linked List from an array
func CreateLinkedListFromArray(values []string) *LinkedList {
	if len(values) > 0 {
		linkedList := CreateLinkedList(values[0])
		for _, value := range values[1:] {
			_, _ = linkedList.Append(value)
		}
		return linkedList
	}
	return nil
}

//Next Used to iterate over the Linked List
func (linkedList *LinkedList) Next() (*Node, bool) {
	if !linkedList.iterating {
		linkedList.iterating = true
		linkedList.Node = linkedList.Head
		return linkedList.Node, linkedList.iterating
	}
	if linkedList.Node != nil && linkedList.Node.Next != nil {
		linkedList.Node = linkedList.Node.Next
		linkedList.iterating = true
		return linkedList.Node, linkedList.iterating
	}
	linkedList.iterating = false
	return nil, linkedList.iterating
}

//Count Get the total number of elements in the Linked List
func (linkedList *LinkedList) Count() int {
	if linkedList.Head == nil || linkedList.GetLast() == nil || linkedList.Length == 0 {
		return 0
	}
	return linkedList.Length
}

//Append Add a new node to the end of the list
func (linkedList *LinkedList) Append(value string) (*Node, error) {
	return linkedList.appendNode(createLinkedListNode(value))
}

//NodeAt Get the node at the an index
func (linkedList *LinkedList) NodeAt(index int) (*Node, error) {
	if linkedList.Head == nil || index < 0 || index > linkedList.Length {
		return nil, fmt.Errorf("error while trying to access node at %d", index)
	}
	var element = linkedList.Head
	for i := 0; i < index; i++ {
		element = element.Next
	}
	return element, nil
}

//Insert Insert a node at a particular index
func (linkedList *LinkedList) Insert(value string, index int) (*Node, error) {
	return linkedList.insertNode(createLinkedListNode(value), index)
}

//Remove a node from a particular index
func (linkedList *LinkedList) Remove(index int) (string, error) {
	currentNode, err := linkedList.NodeAt(index)
	if err != nil {
		return "", err
	}
	node, err := linkedList.removeNode(currentNode)
	if err != nil {
		return "", err
	}
	linkedList.Length--
	if node != nil {
		return node.Value, nil
	}
	return "", nil
}

//RemoveAll Clear all the nodes from the list
func (linkedList *LinkedList) RemoveAll() {
	linkedList.Length = 0
	linkedList.Head = nil
}

//RemoveLast remove the last node from the list
func (linkedList *LinkedList) RemoveLast() (string, error) {
	return linkedList.Remove(linkedList.Length - 1)
}

//Reverse Reverse all the elements of the list such that the Head becomes Last and vice versa
func (linkedList *LinkedList) Reverse() *LinkedList {
	var currentNode = linkedList.Head
	var previousNode *Node
	var nextNode *Node

	for currentNode != nil {
		nextNode = currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}

	linkedList.Head = previousNode
	return linkedList
}

//ToArray Convert the Linked List to an array
func (linkedList *LinkedList) ToArray() ([]string, error) {
	if linkedList.Head == nil || linkedList.Length == 0 {
		return nil, fmt.Errorf("no data to create array")
	}
	var linkedListArray []string
	var element = linkedList.Head
	for i := 0; i < linkedList.Length; i++ {
		linkedListArray = append(linkedListArray, element.Value)
		element = element.Next
	}
	return linkedListArray, nil
}
