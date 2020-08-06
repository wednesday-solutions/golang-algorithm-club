package binarytree

//CreateNode Create a new Node
func CreateNode(Data string, leftNode *BinaryNode, rightNode *BinaryNode) *BinaryNode {
	return &BinaryNode{
		Data:  Data,
		Left:  leftNode,
		Right: rightNode,
	}
}

//BinaryNode Node in a the Binary Tree
type BinaryNode struct {
	BinaryTree
	Data  string
	Left  *BinaryNode
	Right *BinaryNode
}

//BinaryTree BinaryTree interface
type BinaryTree interface {
	TraversePreOrder(func(string))
	TraversePostOrder(func(string))
	TraverseInOrder(func(string))
	IsLeaf() bool
	Count() int64
}

//IsLeaf Check if the current node is a leaf node
func (node *BinaryNode) IsLeaf() bool {
	return node != nil && node.Left == nil && node.Right == nil
}

//TraversePreOrder Traverse the binary tree using PreOrder Traversal
func (node *BinaryNode) TraversePreOrder(fn func(string)) {
	if node == nil {
		return
	}
	fn(node.Data)
	node.Left.TraversePreOrder(fn)
	node.Right.TraversePreOrder(fn)
}

//TraversePostOrder Traverse the binary tree using PostOrder Traversal
func (node *BinaryNode) TraversePostOrder(fn func(string)) {
	if node == nil {
		return
	}
	node.Left.TraversePostOrder(fn)
	node.Right.TraversePostOrder(fn)
	fn(node.Data)
}

//TraverseInOrder Traverse the binary tree using InOrder Traversal
func (node *BinaryNode) TraverseInOrder(fn func(string)) {
	if node == nil {
		return
	}
	node.Left.TraverseInOrder(fn)
	fn(node.Data)
	node.Right.TraverseInOrder(fn)
}

//Count Get a count of the items in the Binary Tree
func (node *BinaryNode) Count() int64 {
	if node == nil {
		return 0
	}
	if node.IsLeaf() {
		return 1
	}
	var leftNodeCount int64 = 0
	var rightNodeCount int64 = 0

	if node.Left != nil {
		leftNodeCount = node.Left.Count()
	}
	if node.Right != nil {
		rightNodeCount = node.Right.Count()
	}
	return leftNodeCount + 1 + rightNodeCount
}
