package tree

//rbtree implements a red black tree

//RedBlackTree is a struct that contains the first node of the red-black tree, and via which methods
//on the tree are defined
type RedBlackTree struct {
	first *Node
}

//Node is a leaf on the tree, or part of a branch
type Node struct {
}
