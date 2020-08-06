package binarytree

import (
	"os"
	"strings"
	"testing"
)

var binaryTree *BinaryNode

func TestMain(m *testing.M) {
	createBinaryTree()
	code := m.Run()
	os.Exit(code)
}
func createBinaryTree() {
	node5 := CreateNode("5", nil, nil)
	nodeA := CreateNode("a", nil, nil)
	node10 := CreateNode("10", nil, nil)
	node4 := CreateNode("4", nil, nil)
	node3 := CreateNode("3", nil, nil)
	nodeB := CreateNode("b", nil, nil)

	aMinus10 := CreateNode("-", nodeA, node10)
	timesLeft := CreateNode("*", node5, aMinus10)

	minus4 := CreateNode("-", nil, node4)
	divide3AndB := CreateNode("/", node3, nodeB)
	timesRight := CreateNode("*", minus4, divide3AndB)

	binaryTree = CreateNode("+", timesLeft, timesRight)
}
func TestCreateNode(t *testing.T) {
	// tests for leaf nodes
	leafNode := CreateNode("A", nil, nil)
	if leafNode.Right != nil {
		t.Errorf("leafNode.Right should be nil")
	}
	if leafNode.Left != nil {
		t.Errorf("leafNode.Left should be nil")
	}

	// tests for node with a left child
	leftChildNode := CreateNode("B", leafNode, nil)
	if leftChildNode.Left == nil {
		t.Errorf("leftChildNode.Left should be initialised")
	}
	if leftChildNode.Right != nil {
		t.Errorf("leftChildNode.Right should be nil")
	}

	// tests for node with a right child
	rightChildNode := CreateNode("C", nil, leafNode)
	if rightChildNode.Right == nil {
		t.Errorf("rightChildNode.Left should be initialised")
	}
	if rightChildNode.Left != nil {
		t.Errorf("rightChildNode.Right should be nil")
	}
}

func TestBinaryNode_IsLeaf(t *testing.T) {
	leafNode := CreateNode("A", nil, nil)
	if !leafNode.IsLeaf() {
		t.Errorf("leafNode should be a leaf node")
	}

	leftChildNode := CreateNode("B", leafNode, nil)
	if leftChildNode.IsLeaf() {
		t.Errorf("leftChildNode should not be a leaf node")
	}

	rightChildNode := CreateNode("B", nil, leafNode)
	if rightChildNode.IsLeaf() {
		t.Errorf("rightChildNode should not be a leaf node")
	}
}

func TestBinaryNode_Count(t *testing.T) {
	if binaryTree.Count() != 12 {
		t.Errorf("Count should be 12")
	}
}

func TestBinaryNode_TraversePreOrder(t *testing.T) {
	expectedResult := "+ * 5 - a 10 * - 4 / 3 b"

	var results []string
	binaryTree.TraversePreOrder(func(data string) {
		results = append(results, data)
	})
	resultString := strings.Join(results, " ")
	if resultString != expectedResult {
		t.Errorf("result should be %s", expectedResult)
	}
}

func TestBinaryNode_TraverseInOrder(t *testing.T) {
	expectedResult := "5 * a - 10 + - 4 * 3 / b"

	var results []string
	binaryTree.TraverseInOrder(func(data string) {
		results = append(results, data)
	})
	resultString := strings.Join(results, " ")
	if resultString != expectedResult {
		t.Errorf("result should be %s", expectedResult)
	}
}

func TestBinaryNode_TraversePostOrder(t *testing.T) {
	expectedResult := "5 a 10 - * 4 - 3 b / * +"
	var results []string
	binaryTree.TraversePostOrder(func(data string) {
		results = append(results, data)
	})
	resultString := strings.Join(results, " ")
	if resultString != expectedResult {
		t.Errorf("result should be %s", expectedResult)
	}
}
