package binarysearchtree

import "fmt"

type (
	// BinarySearchTree implements Insert and Search methods for a BinaryNode
	BinarySearchTree interface {
		Insert(data)
		Search(int) (string, error)
	}
	// BinaryNode a Binary node with data, and left and right child binary node pointers
	BinaryNode struct {
		data
		leftChild  *BinaryNode
		rightChild *BinaryNode
	}
	data struct {
		key  int
		name string
	}
)

// Search for a name in the Binary tree
func (node *BinaryNode) Search(searchKey int) (name string, err error) {
	err = nil
	name = ""
	switch {
	case node.data.key == searchKey:
		name = node.data.name
		return name, err
	case node.data.key < searchKey:
		if node.rightChild.data.name != "" {
			return node.rightChild.Search(searchKey)
		}
		err = fmt.Errorf("Not found")
		return name, err
	case node.data.key > searchKey:
		if node.leftChild.data.name != "" {
			return node.leftChild.Search(searchKey)
		}
		err = fmt.Errorf("Not found")
		return name, err
	}
	return
}

// Insert data into a node according to key
func (node *BinaryNode) Insert(input data) {
	switch {
	case node == nil || node.data.name == "":
		// node is empty
		node.data = input
	case node.data.key > input.key:
		// key less than node, save data in the left branch
		if node.leftChild == nil {
			// empty left child
			node.leftChild = &BinaryNode{input, nil, nil}
		} else {
			node.leftChild.Insert(input)
		}
	case node.data.key < input.key:
		// key greater than node key, save data in the right branch
		if node.rightChild == nil {
			// empty left child
			node.rightChild = &BinaryNode{input, nil, nil}
		} else {
			node.rightChild.Insert(input)
		}
	}
}
