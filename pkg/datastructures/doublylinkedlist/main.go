package doublylinkedlist

import "fmt"

type (
	//Node Doubly Linked List node
	Node struct {
		Value    string
		Previous *Node
		Next     *Node
	}
	//DoublyLinkedList Doubly Linked List struct
	DoublyLinkedList struct {
		DLLInterface
		Head      *Node
		Last      *Node
		Node      *Node
		Length    int
		iterating bool
	}
	//DLLInterface Interface for Doubly Linked List
	DLLInterface interface {
		Append(value string) (*Node, error)
		Count() int
		Insert(value string, index int) (*Node, error)
		Next() (*Node, bool)
		NodeAt(index int) (*Node, error)
		Remove(index int) (string, error)
		RemoveAll()
		RemoveLast() (string, error)
		Reverse() *DoublyLinkedList
		ToArray() ([]string, error)
	}
)

func createDoublyLinkedListNode(value string) *Node {
	return &Node{
		Value:    value,
		Previous: nil,
		Next:     nil,
	}
}

func (doublyLinkedList *DoublyLinkedList) appendNode(node *Node) (*Node, error) {
	if doublyLinkedList.Head == nil && doublyLinkedList.Last == nil {
		doublyLinkedList.Head = node
		doublyLinkedList.Last = node
	} else {
		node.Previous = doublyLinkedList.Last
		doublyLinkedList.Last.Next = node
		doublyLinkedList.Last = node
	}

	doublyLinkedList.Length++
	return node, nil
}
func (doublyLinkedList *DoublyLinkedList) insertNode(node *Node, index int) (*Node, error) {
	if doublyLinkedList.Head == nil {
		doublyLinkedList.Head = node
		doublyLinkedList.Last = node
	} else if index < 0 || index > doublyLinkedList.Length {
		return nil, fmt.Errorf("invalid index")
	} else if index == 0 {
		node.Next = doublyLinkedList.Head
		doublyLinkedList.Head.Previous = node
		doublyLinkedList.Head = node
	} else {
		nodeToReplace, err := doublyLinkedList.NodeAt(index)
		if err != nil {
			return nil, err
		}
		previousNode := nodeToReplace.Previous
		previousNode.Next = node
		node.Previous = previousNode
		node.Next = nodeToReplace
		nodeToReplace.Previous = node
	}
	doublyLinkedList.Length++
	return node, nil
}
func (doublyLinkedList *DoublyLinkedList) removeNode(node *Node) (*Node, error) {

	if doublyLinkedList.Length == 0 || (node.Previous == nil && node.Next == nil) {
		return nil, fmt.Errorf("invalid node")
	}

	previousElement := node.Previous
	nextElement := node.Next

	if previousElement != nil {
		previousElement.Next = nextElement
	} else {
		doublyLinkedList.Head = nextElement
	}

	if nextElement != nil {
		nextElement.Previous = previousElement
	} else {
		doublyLinkedList.Last = previousElement
	}
	return node, nil
}

//CreateDoublyLinkedList Create a new Doubly Linked List
func CreateDoublyLinkedList(value string) *DoublyLinkedList {
	node := createDoublyLinkedListNode(value)
	return &DoublyLinkedList{
		Head:   node,
		Last:   node,
		Length: 1,
		Node:   nil,
	}
}

//Next Used to iterate over the Doubly Linked List
func (doublyLinkedList *DoublyLinkedList) Next() (*Node, bool) {
	if !doublyLinkedList.iterating {
		doublyLinkedList.iterating = true
		doublyLinkedList.Node = doublyLinkedList.Head
		return doublyLinkedList.Node, doublyLinkedList.iterating
	}
	if doublyLinkedList.Node != nil && doublyLinkedList.Node.Next != nil {
		doublyLinkedList.Node = doublyLinkedList.Node.Next
		doublyLinkedList.iterating = true
		return doublyLinkedList.Node, doublyLinkedList.iterating
	}
	doublyLinkedList.iterating = false
	return nil, doublyLinkedList.iterating
}

//Count Get the total number of elements in the Doubly Linked List
func (doublyLinkedList *DoublyLinkedList) Count() int {
	if doublyLinkedList.Head == nil || doublyLinkedList.Last == nil || doublyLinkedList.Length == 0 {
		return 0
	}
	return doublyLinkedList.Length
}

//Append Add a new node to the end of the list
func (doublyLinkedList *DoublyLinkedList) Append(value string) (*Node, error) {
	return doublyLinkedList.appendNode(createDoublyLinkedListNode(value))
}

//NodeAt Get the node at the an index
func (doublyLinkedList *DoublyLinkedList) NodeAt(index int) (*Node, error) {
	if doublyLinkedList.Head == nil || index < 0 || index > doublyLinkedList.Length {
		return nil, fmt.Errorf("error while trying to access node at %d", index)
	}
	var element = doublyLinkedList.Head
	for i := 0; i < index; i++ {
		element = element.Next
	}
	return element, nil
}

//Insert Insert a node at a particular index
func (doublyLinkedList *DoublyLinkedList) Insert(value string, index int) (*Node, error) {
	return doublyLinkedList.insertNode(createDoublyLinkedListNode(value), index)
}

//Remove a node from a particular index
func (doublyLinkedList *DoublyLinkedList) Remove(index int) (string, error) {
	currentNode, err := doublyLinkedList.NodeAt(index)
	if err != nil {
		return "", err
	}
	node, err := doublyLinkedList.removeNode(currentNode)
	if err != nil {
		return "", err
	}
	doublyLinkedList.Length--
	if node != nil {
		return node.Value, nil
	}
	return "", nil
}

//RemoveAll Clear all the nodes from the list
func (doublyLinkedList *DoublyLinkedList) RemoveAll() {
	doublyLinkedList.Length = 0
	doublyLinkedList.Head = nil
	doublyLinkedList.Last = nil
}

//RemoveLast remove the last node from the list
func (doublyLinkedList *DoublyLinkedList) RemoveLast() (string, error) {
	return doublyLinkedList.Remove(doublyLinkedList.Length - 1)
}

//Reverse Reverse all the elements of the list such that the Head becomes Last and vice versa
func (doublyLinkedList *DoublyLinkedList) Reverse() *DoublyLinkedList {
	var currentNode = doublyLinkedList.Head
	var previousNode *Node
	var nextNode *Node

	for currentNode != nil {
		nextNode = currentNode.Next
		previousNode = currentNode.Previous

		currentNode.Next = previousNode
		currentNode.Previous = nextNode

		previousNode = currentNode
		currentNode = nextNode
	}

	doublyLinkedList.Last = doublyLinkedList.Head
	doublyLinkedList.Head = previousNode
	return doublyLinkedList
}

//ToArray Convert the Doubly Linked List to an array
func (doublyLinkedList *DoublyLinkedList) ToArray() ([]string, error) {
	if doublyLinkedList.Head == nil || doublyLinkedList.Last == nil || doublyLinkedList.Length == 0 {
		return nil, fmt.Errorf("no data to create array")
	}
	var doublyLinkedListArray []string
	var element = doublyLinkedList.Head
	for i := 0; i < doublyLinkedList.Length; i++ {
		doublyLinkedListArray = append(doublyLinkedListArray, element.Value)
		element = element.Next
	}
	return doublyLinkedListArray, nil
}
