package redblacktree

import (
	"bytes"
	"errors"
)

//rbtree implements a red black tree

var errorNullNodeRotate = errors.New("Error: Cannot rotate null node")
var errorNullChildren = errors.New("Error: Node contains no children") //might not actually use this
var errorNullChild = errors.New("Error: Child not found")

//RedBlackTree is a struct that contains the first node of the red-black tree, and via which methods
//on the tree are defined
type RedBlackTree struct {
	first *Node
}

//Node is a leaf on the tree, or part of a branch
type Node struct {
	data       []byte
	parent     *Node
	leftChild  *Node
	rightChild *Node
}

func newNode(data []byte) *Node {
	return &Node{
		data: data,
	}
}

//Put takes a byte slice as an argument, creates a node, and adds it to the tree
func (b *RedBlackTree) Put(data []byte) {
	//create node
	node := newNode(data)
	if b.first == nil {
		b.first = node
		return
	}
	current := b.first
	for { //no conditions because we'll return within the loop
		//compare node to current node
		res := bytes.Compare(node.data, current.data)
		if res < 0 { //if new node is smaller than current node
			if current.leftChild == nil {
				current.leftChild = node
				node.parent = current
				break
			} else {
				current = current.leftChild
			}
		} else if res > 0 { //if new node is greater than current node
			if current.rightChild == nil {
				current.rightChild = node
				node.parent = current
				break
			} else {
				current = current.rightChild
			}
		} else {
			//if they are equal, entry is already in tree
			break
		}
	}
}

//Search takes a slice of bytes as an argument and searches the tree for that data, returning
//the node containing the data, or nil if the data is not found in the tree
func (b *RedBlackTree) Search(data []byte, n *Node) *Node {
	if n == nil {
		return n
	}
	res := bytes.Compare(data, n.data)
	if res < 0 {
		return b.Search(data, n.leftChild)
	} else if res > 0 {
		return b.Search(data, n.rightChild)
	} else {
		return n
	}
}

//RotateLeft switches a parent node and its right child
func (b *RedBlackTree) rotateLeft(n *Node) error {
	if n == nil {
		return errorNullNodeRotate
	}
	if n.rightChild == nil {
		return errorNullChild
	}

	//make temp node from right child
	res := n.rightChild
	//make res's left child the right child of node
	n.rightChild = res.leftChild
	//make res's parent the node's parent
	res.parent = n.parent
	//if res is head of tree, mark in tree. Otherwise, make child of parent node
	if res.parent == nil {
		b.first = res
	} else {
		if n.parent.leftChild == n {
			n.parent.leftChild = res
		} else {
			n.parent.rightChild = res
		}
	}
	//make node's parent res
	n.parent = res

	return nil
}

//Rotateright switches a parent node and its left child
func (b *RedBlackTree) rotateRight(n *Node) error {
	if n == nil {
		return errorNullNodeRotate
	}
	if n.leftChild == nil {
		return errorNullChild
	}

	//make temp node from left child
	res := n.leftChild
	//make res's right child the left child of node
	n.leftChild = res.rightChild
	//make res's parent the node's parent
	res.parent = n.parent
	//if res is head of tree, mark in tree. Otherwise, make child of parent node
	if res.parent == nil {
		b.first = res
	} else {
		if n.parent.leftChild == n {
			n.parent.leftChild = res
		} else {
			n.parent.rightChild = res
		}
	}
	//make node's parent res
	n.parent = res

	return nil
}
