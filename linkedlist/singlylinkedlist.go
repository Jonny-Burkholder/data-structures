package linkedlist

import "fmt"

//this is the island of misfit structures

//LinkedList is a single linked list. This is based more on the methods and items it contains,
//rather than on the structure itself
type LinkedList struct {
	Length int
	Head   *Item
	Tail   *Item
}

//Item is an entry in a single-linked list
type Item struct {
	Data interface{}
	Next *Item
}

//NewLinkedList takes no parameters and returns an empty singly linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

//NewItem takes data of empty interface type and returns a pointer to type Item
func (l *LinkedList) newItem(data interface{}) *Item {
	return &Item{
		Data: data,
	}
}

//Next takes an input item and returns the item directly after it
func (l *LinkedList) Next(item *Item) (*Item, error) {
	if item == l.Tail {
		return &Item{}, fmt.Errorf("No more items in list")
	}

	return item.Next, nil
}

//Search takes data of type interface{} as input, and iterates through all items in a list until
//a match is found, or an error if no match is found
func (l *LinkedList) Search(data interface{}) (*Item, error) {
	current := l.Head
	for i := 0; i < l.Length-1; i++ {
		if current.Data == data {
			return current, nil
		}
		current = current.Next
	}
	return &Item{}, fmt.Errorf("Error: no match found for %s", data)
}

//Put places a new item at the end of the list
func (l *LinkedList) Put(data interface{}) {
	item := l.newItem(data)
	if l.Length == 0 {
		l.Head = item
		l.Length++
		return
	}
	current := l.Head
	for i := 0; i < l.Length-1; i++ {
		current = current.Next
	}
	current.Next = item
	l.Length++
}

//Push places a new item at the head of the list
func (l *LinkedList) Push(data interface{}) {
	item := l.newItem(data)
	if l.Length == 0 {
		l.Head = item
		l.Length++
	} else if l.Length == 1 {
		l.Head.Next = item
		l.Length++
	} else {
		current := l.Head
		for i := 0; i < l.Length-1; i++ { //should get us to the last item. I can't brain, I'll have to test it
			current = current.Next
		}
		current.Next = item
		l.Length++
	}
}

//Pop removes and returns the last item in the list
func (l *LinkedList) Pop() *Item {
	current := l.Head
	for i := 0 + 1; i < l.Length-1; i++ { //we actually want to get to the penultimate item
		current = current.Next
	}
	res := current.Next
	current.Next = &Item{}
	l.Length--
	return res
}

//Delete removes an item from the list, or returns an error if the item is not found
func (l *LinkedList) Delete(item *Item) error {
	current := l.Head
	if current == item {
		l.Head = current.Next
		l.Length--
		return nil
	}
	for i := 0; i < l.Length-1; i++ {
		previous := current
		current = current.Next
		if current == item {
			if current == l.Tail {
				l.Tail = previous
				previous.Next = &Item{}
				l.Length--
				return nil
			}
			previous.Next = current.Next
			l.Length--
			return nil
		}
	}
	return fmt.Errorf("Error: item not found")
}

//DeleteIndex takes an int as input and removes the item at that index from the list, where applicaple
func (l *LinkedList) DeleteIndex(index int) error {
	if index < 0 {
		return fmt.Errorf("Try a positive number, joker")
	}
	if index == 0 {
		l.Head = l.Head.Next
		l.Length--
		return nil
	}

	if index >= l.Length {
		return fmt.Errorf("Error: index %v out of range", index)
	}

	previous := l.Head
	current := l.Head.Next

	for i := 0 + 1; i == index; i++ { //I know it's a little confusing, I just don't want my index to start at 1
		previous = current
		current = current.Next
	}
	if current == l.Tail {
		l.Tail = previous
		previous.Next = &Item{}
		l.Length--
		return nil
	}
	previous.Next = current.Next
	l.Length--
	return nil
}
