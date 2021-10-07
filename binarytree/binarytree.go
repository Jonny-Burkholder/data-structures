package binarytree

import (
	"bytes"
	"errors"
)

var errorNullNodeRotate = errors.New("Error: Cannot rotate null node")
var errorNullChildren = errors.New("Error: Node contains no children")
var errorNullChild = errors.New("Error: Child not found")

//BinaryTree has a root node pointer and implements basic BinaryTree functions
type BinaryTree struct {
	First *Node
}

//Node is an entry in the tree
type Node struct {
	Data       []byte
	LeftChild  *Node
	RightChild *Node
}

func newNode(data []byte) *Node {
	return &Node{
		Data: data,
	}
}

//NewBinaryTree creates and returns a BinaryTree
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

//AddEntry takes a byte slice as an argument, creates a node, and adds it to the tree
func (b *BinaryTree) AddEntry(data []byte) {
	//create node
	node := newNode(data)
	current := b.First
	for { //no conditions because we'll return within the loop
		//compare node to current node
		res := bytes.Compare(node.Data, current.Data)
		if res < 0 { //if new node is smaller than current node
			if current.LeftChild == nil {
				current.LeftChild = node
				break
			} else {
				current = current.LeftChild
			}
		} else if res > 0 { //if new node is greater than current node
			if current.RightChild == nil {
				current.RightChild = node
				break
			} else {
				current = current.RightChild
			}
		} else {
			break
		}
	}
}

//Search takes a slice of bytes as an argument and searches the tree for that data, returning
//the node containing the data, or nil if the data is not found in the tree
func (b *BinaryTree) Search(data []byte, n *Node) *Node {
	if n == nil {
		return n
	}
	res := bytes.Compare(data, n.Data)
	if res < 0 {
		return b.Search(data, n.LeftChild)
	} else if res > 0 {
		return b.Search(data, n.RightChild)
	} else {
		return n
	}
}
