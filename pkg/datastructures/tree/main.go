package tree

//Node Node of a Tree
type Node struct {
	Tree
	Value    string
	Parent   *Node
	Children []*Node
}

//Tree Tree interface
type Tree interface {
	AddChild(*Node)
	Search(value string) *Node
}

//NewNode Create a New Node
func NewNode(value string) *Node {
	return &Node{
		Value:    value,
		Parent:   nil,
		Children: nil,
	}
}

//AddChild Add a child to the node
func (node *Node) AddChild(n *Node) {
	node.Children = append(node.Children, n)
	n.Parent = node
}

//Search Look for a child with the same value as the passed in value
func (node *Node) Search(value string) *Node {
	if value == node.Value {
		return node
	}
	var searchNode *Node
	for _, child := range node.Children {
		searchNode = child.Search(value)
		if node != nil {
			break
		}
	}
	return searchNode
}
