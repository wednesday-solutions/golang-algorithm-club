package binarysearchtree

import (
	"github.com/wednesday-solutions/golang-algorithm-club/pkg/utl"
	"testing"
)

var binaryTree = BinaryNode{}

func TestMain(m *testing.M) {
	utl.SetupTests(m, initBinarySearchTree)
}

func initBinarySearchTree() {
	binaryTree.Insert(data{key: 0, name: "NODE ZERO"})
	binaryTree.Insert(data{key: -1, name: "LEFT CHILD"})
	binaryTree.Insert(data{key: 1, name: "RIGHT CHILD"})
}

func TestInsert(t *testing.T) {
	if binaryTree.data.name != "NODE ZERO" {
		t.Errorf("NODE ZERO should be inserted in the root node")
	}
	if binaryTree.leftChild.data.name != "LEFT CHILD" {
		t.Errorf("LEFT CHILD should be inserted in the left child of the root node")
	}
	if binaryTree.rightChild.data.name != "RIGHT CHILD" {
		t.Errorf("RIGHT CHILD should be inserted in the right child of the root node")
	}
}

func TestSearch(t *testing.T) {
	keyZero, err := binaryTree.Search(0)
	if keyZero != "NODE ZERO" && err != nil {
		t.Errorf("Search for key 0 should return NODE ZERO, with no errors")
	}
	keyOne, err := binaryTree.Search(1)
	if keyOne != "RIGHT CHILD" && err != nil {
		t.Errorf("Search for key 1 should return RIGHT CHILD, with no errors")
	}
	keyMinusOne, err := binaryTree.Search(-1)
	if keyMinusOne != "LEFT CHILD" && err != nil {
		t.Errorf("Search for key -1 should return LEFT CHILD, with no errors")
	}
}
