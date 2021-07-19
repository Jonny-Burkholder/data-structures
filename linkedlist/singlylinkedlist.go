package linkedlist

import "fmt"

//this is the island of misfit structures

//LinkedList is a single linked list. This is based more on the methods and items it contains,
//rather than on the structure itself
type linkedList struct {
	Length int
	Head   *item
	Tail   *item
}

//Item is an entry in a single-linked list
type item struct {
	data interface{}
	next *item
}

//NewLinkedList takes no parameters and returns an empty singly linked list
func NewLinkedList() *linkedList {
	return &linkedList{}
}

//NewItem takes data of empty interface type and returns a pointer to type Item
func (l *linkedList) newItem(data interface{}) *item {
	return &item{
		data: data,
	}
}

//Next takes an input item and returns the item directly after it
func (l *linkedList) Next(it *item) (*item, error) {
	if it == l.Tail {
		return &item{}, fmt.Errorf("No more items in list")
	}

	return it.next, nil
}

//Search takes data of type interface{} as input, and iterates through all items in a list until
//a match is found, or an error if no match is found
func (l *linkedList) Search(data interface{}) (*item, error) {
	current := l.Head
	for i := 0; i < l.Length-1; i++ {
		if current.data == data {
			return current, nil
		}
		current = current.next
	}
	return &item{}, fmt.Errorf("Error: no match found for %s", data)
}

//Push places a new item at the end of the list
func (l *linkedList) Push(data interface{}) {
	item := l.newItem(data)
	if l.Length == 0 {
		l.Head = item
		l.Tail = item
		l.Length++
	} else if l.Length == 1 {
		l.Head.next = item
		l.Tail = item
		l.Length++
	} else {
		current := l.Head
		for i := 0; i < l.Length-1; i++ { //should get us to the last item. I can't brain, I'll have to test it
			current = current.next
		}
		current.next = item
		l.Tail = item
		l.Length++
	}
}

//Pop removes and returns the last item in the list
func (l *linkedList) Pop() *item {
	current := l.Head
	for i := 0 + 1; i < l.Length-1; i++ { //we actually want to get to the penultimate item
		current = current.next
	}
	res := current.next
	current.next = &item{}
	l.Tail = current
	l.Length--
	return res
}

//Unshift places a new item at the head of the list
func (l *linkedList) Unshift(data interface{}) {
	item := l.newItem(data)
	if l.Length == 0 {
		l.Head = item
		l.Tail = item
	} else {
		item.next = l.Head
		l.Head = item
	}
	l.Length++
}

//Shift removes an item from the beginning of the list
func (l *linkedList) Shift() (*item, error) {
	if l.Length < 1 {
		return &item{}, fmt.Errorf("Error: no items in list")
	}

	it := l.Head
	l.Length--

	if l.Length == 0 {
		l.Head = &item{}
		l.Tail = &item{}
	} else if l.Length == 1 {
		l.Head = l.Tail
	} else {
		l.Head = l.Head.next
	}
	return it, nil
}

//Delete removes an item from the list, or returns an error if the item is not found
func (l *linkedList) Delete(it *item) error {
	current := l.Head
	if current == it {
		l.Head = current.next
		l.Length--
		return nil
	}
	for i := 0; i < l.Length-1; i++ {
		previous := current
		current = current.next
		if current == it {
			if current == l.Tail {
				l.Tail = previous
				previous.next = &item{}
				l.Length--
				return nil
			}
			previous.next = current.next
			l.Length--
			return nil
		}
	}
	return fmt.Errorf("Error: item not found")
}

//DeleteIndex takes an int as input and removes the item at that index from the list, where applicaple
func (l *linkedList) DeleteIndex(index int) error {
	if index < 0 {
		return fmt.Errorf("Try a positive number, joker")
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
		previous.next = &item{}
		l.Length--
		return nil
	}
	previous.next = current.next
	l.Length--
	return nil
}
