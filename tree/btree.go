package tree

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
	//node := newNode(data)
}
