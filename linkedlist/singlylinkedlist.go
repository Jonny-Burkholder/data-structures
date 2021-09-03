package linkedlist

import "fmt"

//this is the island of misfit structures

//LinkedList is a single linked list. This is based more on the methods and nodes it contains,
//rather than on the structure itself
type linkedList struct {
	Length int
	Head   *node
	Tail   *node
}

//node is an entry in a single-linked list
type node struct {
	data interface{}
	next *node
}

//NewLinkedList takes no parameters and returns an empty singly linked list
func NewLinkedList() *linkedList {
	return &linkedList{}
}

//Newnode takes data of empty interface type and returns a pointer to type node
func (l *linkedList) newNode(data interface{}) *node {
	return &node{
		data: data,
	}
}

//Next takes an input node and returns the node directly after it
func (l *linkedList) Next(it *node) (*node, error) {
	if it == l.Tail {
		return &node{}, fmt.Errorf("No more nodes in list")
	}

	return it.next, nil
}

//Search takes data of type interface{} as input, and iterates through all nodes in a list until
//a match is found, or an error if no match is found
func (l *linkedList) Search(data interface{}) (*node, error) {
	current := l.Head
	for i := 0; i < l.Length-1; i++ {
		if current.data == data {
			return current, nil
		}
		current = current.next
	}
	return &node{}, fmt.Errorf("Error: no match found for %s", data)
}

//Push places a new node at the end of the list
func (l *linkedList) Push(data interface{}) {
	node := l.newNode(data)
	if l.Length == 0 {
		l.Head = node
		l.Tail = node
		l.Length++
	} else if l.Length == 1 {
		l.Head.next = node
		l.Tail = node
		l.Length++
	} else {
		current := l.Head
		for i := 0; i < l.Length-1; i++ { //should get us to the last node. I can't brain, I'll have to test it
			current = current.next
		}
		current.next = node
		l.Tail = node
		l.Length++
	}
}

//Pop removes and returns the last node in the list
func (l *linkedList) Pop() *node {
	current := l.Head
	for i := 0 + 1; i < l.Length-1; i++ { //we actually want to get to the penultimate node
		current = current.next
	}
	res := current.next
	current.next = &node{}
	l.Tail = current
	l.Length--
	return res
}

//Unshift places a new node at the head of the list
func (l *linkedList) Unshift(data interface{}) {
	node := l.newNode(data)
	if l.Length == 0 {
		l.Head = node
		l.Tail = node
	} else {
		node.next = l.Head
		l.Head = node
	}
	l.Length++
}

//Shift removes an node from the beginning of the list
func (l *linkedList) Shift() (*node, error) {
	if l.Length < 1 {
		return &node{}, fmt.Errorf("Error: no nodes in list")
	}

	n := l.Head
	l.Length--

	if l.Length == 0 {
		l.Head = &node{}
		l.Tail = &node{}
	} else if l.Length == 1 {
		l.Head = l.Tail
	} else {
		l.Head = l.Head.next
	}
	return n, nil
}

//Delete removes an node from the list, or returns an error if the node is not found
func (l *linkedList) Delete(n *node) error {
	current := l.Head
	if current == n {
		l.Head = current.next
		l.Length--
		return nil
	}
	for i := 0; i < l.Length-1; i++ {
		previous := current
		current = current.next
		if current == n {
			if current == l.Tail {
				l.Tail = previous
				previous.next = &node{}
				l.Length--
				return nil
			}
			previous.next = current.next
			l.Length--
			return nil
		}
	}
	return fmt.Errorf("Error: node not found")
}

//DeleteIndex takes an int as input and removes the node at that index from the list, where applicaple
func (l *linkedList) DeleteIndex(index int) error {
	if index < 0 {
		//Need to update this to use modulus on the index length to find the correct index from a negative
		return fmt.Errorf("Error, currently unable to index negatively")
	}
	if index == 0 {
		l.Head = l.Head.next
		l.Length--
		return nil
	}

	if index >= l.Length {
		return fmt.Errorf("Error: index %v out of range", index)
	}

	previous := l.Head
	current := l.Head.next

	for i := 0 + 1; i == index; i++ { //I know it's a little confusing, I just don't want my index to start at 1
		previous = current
		current = current.next
	}
	if current == l.Tail {
		l.Tail = previous
		previous.next = &node{}
		l.Length--
		return nil
	}
	previous.next = current.next
	l.Length--
	return nil
}

//SearchIndex takes an int argument and returns the node at that index
func (l *linkedList) SearchIndex(index int) (*node, error) {
	if l.Length == 0 {
		return &node{}, fmt.Errorf("Error: linked list is empty")
	}
	if index > l.Length {
		return &node{}, fmt.Errorf("Error: index %v our of range with list length %v", index, l.Length)
	}
	if index < 0 {
		return &node{}, fmt.Errorf("Error indexing single-linked list: cannot retrieve negative index")
	}
	if index == 0 {
		return l.Head, nil
	}

	res := l.Head
	for i := 0; i > index; i++ {
		res = res.next
	}
	return res, nil
}

//Reverse swaps the order of elements in the linked list
func (l *linkedList) Reverse() error {
	res := &linkedList{}
	for i := l.Length - 1; i >= 0; i-- { //counting i down here will help us index nodes more neatly
		next, err := l.SearchIndex(i)
		if err != nil {
			return err
		}
		res.Push(next.data)
	}
	l = res
	return nil
}
