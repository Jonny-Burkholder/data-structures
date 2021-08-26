package tree

import (
	"bytes"
)

//BTree has a root node pointer and implements basic BTree functions
type BTree struct {
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

//NewBTree creates and returns a BTree
func NewBTree() *BTree {
	return &BTree{}
}

//AddEntry takes a byte slice as an argument, creates a node, and adds it to the tree
func (b *BTree) AddEntry(data []byte) {
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
//the node containing the data, or a non-nil error if the data is not found in the tree
func (b *BTree) Search(data []byte, n *Node) *Node {
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
