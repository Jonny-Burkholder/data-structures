package linkedlist

import "fmt"

//DoublyLinkedList is a struct that contains a pointer to a header file, a pointer to a tail file, a
//length integer. Its methods and items allow it to be doubly-linked
type doublyLinkedList struct {
	Length int
	Head   *doubleItem
	Tail   *doubleItem
}

//doubleItem is an item that contains a pointer to both the previous item and the next item in the list,
//as well as data of type interface{}
type doubleItem struct {
	data     interface{}
	previous *doubleItem
	next     *doubleItem
}

//NewDoublyLinkedList returns a pointer to an empty doublyLinkedList
func NewDoublyLinkedList() *doublyLinkedList {
	return &doublyLinkedList{}
}

//newDoubleItem takes data of type interface{} and returns a pointer to a doubleItem with data, but no
//pointers to previous or next
func newDoubleItem(data interface{}) *doubleItem {
	return &doubleItem{
		data: data,
	}
}

//Put places an item at the end of the list
func (l *doublyLinkedList) Put(data interface{}) {
	item := newDoubleItem(data)
	if l.Length == 0 {
		l.Head = item
		l.Tail = item
	} else if l.Length == 1 {
		l.Head.next = item
		l.Tail = item
	} else {
		l.Tail.next = item
		item.previous = l.Tail
		l.Tail = item
	}
	l.Length++
}

//Push places an item at the beginning of the list
func (l *doublyLinkedList) Push(data interface{}) {
	item := newDoubleItem(data)
	if l.Length == 0 {
		l.Head = item
		l.Tail = item
	} else if l.Length == 1 {
		l.Head = item
		l.Head.next = l.Tail
		l.Tail.previous = item
	} else {
		l.Head.previous = item
		item.next = l.Head
		l.Head = item
	}
	l.Length++
}

//Pop removes an item from the end of the list, or returns a non-nil error if there are no items in
//the list
func (l *doublyLinkedList) Pop() (*doubleItem, error) {
	if l.Length < 1 {
		return &doubleItem{}, fmt.Errorf("Error: there are no items in list")
	}

	item := l.Tail

	if l.Length == 1 {
		l.Head = &doubleItem{}
		l.Tail = &doubleItem{}
	} else {
		l.Tail = l.Tail.previous
		l.Tail.next = &doubleItem{}
	}
	l.Length--
	return item, nil
}

//Shift removes an item from the beginning of the list, or returns a non-nil error if list is empty
func (l *doublyLinkedList) Shift() (*doubleItem, error) {
	if l.Length < 1 {
		return &doubleItem{}, fmt.Errorf("Error: there are no items in list")
	}

	item := l.Head

	if l.Length == 1 {
		l.Head = &doubleItem{}
		l.Tail = &doubleItem{}
	} else {
		l.Head = l.Head.next
		l.Head.previous = &doubleItem{}
	}
	l.Length--
	return item, nil
}

//Next takes a pointer to a doubleItem as an input and returns a pointer to the next doubleItem in the
//list. It returns a non-nill error if the item is the final item in the list
func (l *doublyLinkedList) Next(item *doubleItem) (*doubleItem, error) {
	if item != l.Tail {
		return item.next, nil
	}
	return &doubleItem{}, fmt.Errorf("Error: index out of range")
}

//Previous takes a pointer to a doubleItem as an input and returns a pointer to the previous doubleItem
//in the list. It returns a non-nill error if the item is the first item in the list
func (l *doublyLinkedList) Previous(item *doubleItem) (*doubleItem, error) {
	if item != l.Head {
		return item.previous, nil
	}
	return &doubleItem{}, fmt.Errorf("Error: index out of range")
}

//Search takes data of type interface{} and iterates through each item in the list until the item is
//found. If it is not found, then a non-nil error is returned
func (l *doublyLinkedList) Search(data interface{}) (*doubleItem, error) {
	current := l.Head
	for i := 0; i < l.Length-1; i++ {
		if current.data != data {
			current = current.next
		} else {
			return current, nil
		}
	}
	return &doubleItem{}, fmt.Errorf("Data not found in list")
}

//Delete removes an item based on the data it contains, or returns a non-nil error if the item is not found
func (l *doublyLinkedList) Delete(data interface{}) error {
	item, err := l.Search(data)
	if err != nil {
		return err
	}
	item.previous.next = item.next
	item.next.previous = item.previous
	l.Length--
	return nil
}
