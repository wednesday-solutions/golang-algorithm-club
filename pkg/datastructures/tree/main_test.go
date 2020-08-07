package tree

import (
	"github.com/wednesday-solutions/golang-algorithm-club/pkg/utl"
	"testing"
)

var root *Node

func createTree() {
	root = NewNode("100")
	child1 := NewNode("200")
	child2 := NewNode("300")
	child3 := NewNode("400")
	child4 := NewNode("500")
	root.AddChild(child1)
	root.AddChild(child2)
	child1.AddChild(child3)
	child2.AddChild(child4)

}
func TestMain(m *testing.M) {
	utl.SetupTests(m, createTree)
}

func TestNewNode(t *testing.T) {
	value := "100"
	node := NewNode(value)
	if node.Value != value {
		t.Errorf("NewNode(%s) should create a node with alue: %s", value, value)
	}
	if node.Parent != nil {
		t.Errorf("NewNode(%s) should create a node with Parent nil", value)
	}
	if node.Children != nil {
		t.Errorf("NewNode(%s) should create a node with Children nil", value)
	}
}

func TestNode_AddChild(t *testing.T) {
	node := NewNode("999")
	root.AddChild(node)
	if node.Parent != root {
		t.Errorf("the root node should be the parent of the newly added node")
	}
	if root.Children[len(root.Children)-1] != node {
		t.Errorf("the newly added node must be the last child of the root node")
	}
}

func TestNode_Search(t *testing.T) {
	lookedUpNode := root.Search("200")
	if lookedUpNode != root.Children[0] {
		t.Errorf("node with value 200 should be the first child of the root node")
	}
}
