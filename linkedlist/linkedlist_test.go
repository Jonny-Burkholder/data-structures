package linkedlist

import (
	"fmt"
	"testing"
)

func TestSinglePush(t *testing.T) {
	l := NewLinkedList()
	l.Push("one")
	l.Push("two")
	l.Push("three")
	if l.Head.data != "one" {
		t.Errorf("Incorrect item, wanted %s, got %s", "one", l.Head.data)
	}
	if l.Tail.data != "three" {
		t.Errorf("Incorrect item, wanted %s, got %s", "three", l.Tail.data)
	}
	if l.Length != 3 {
		t.Errorf("Incorrect list length, wanted %v, got %v", 3, l.Length)
	}
}

func TestSinglePop(t *testing.T) {
	l := NewLinkedList()
	l.Push("one")
	l.Push("two")
	l.Push("three")

	data := l.Pop().data
	if data != "three" {
		t.Errorf("Incorrect data, wanted %s, got %s", "three", data)
	}

	if l.Tail.data != "two" {
		t.Errorf("Incorrect tail, wanted %s, got %s", "two", l.Tail.data)
	}

	if l.Length != 2 {
		t.Errorf("Incorrect list length, wanted %v, got %v", 2, l.Length)
	}

}

func TestSingleUnshift(t *testing.T) {
	l := NewLinkedList()
	l.Unshift("three")
	l.Unshift("two")
	l.Unshift("one")

	if l.Head.data != "one" {
		t.Errorf("Incorrect head data, wanted %s, got %s", "one", l.Head.data)
	}
	if l.Tail.data != "three" {
		t.Errorf("Incorrect tail data, wanted %s, got %s", "three", l.Tail.data)
	}
	if l.Length != 3 {
		t.Errorf("Incorrect list lenght, wanted %v, got %v", 3, l.Length)
	}
}

func TestSingleShift(t *testing.T) {
	l := NewLinkedList()
	l.Unshift(3)
	l.Unshift(2)
	l.Unshift(1)

	data, _ := l.Shift()

	if data.data != 1 {
		t.Errorf("Incorrect data, wanted %v, got %v", 1, data.data)
	}
	if l.Head.data != 2 {
		t.Errorf("Incorrect head data, wanted %v, got %v", 2, l.Head.data)
	}
	if l.Tail.data != 3 {
		t.Errorf("Incorrect tail data, wanted %v, got %v", 3, l.Tail.data)
	}
}

func TestSingleSearch(t *testing.T) {
	l := NewLinkedList()
	l.Push("A")
	l.Push("B")
	l.Push("b")
	l.Push("c")

	item, err := l.Search("b")

	if err != nil {
		t.Error(err)
	}

	if item.data != "b" {
		t.Errorf("Incorrect item returned. Wanted %s, got %s", "b", item.data)
	}
}

func TestSingleNext(t *testing.T) {
	l := NewLinkedList()
	l.Push("one")
	l.Push(2)
	l.Push(3.0)

	item, _ := l.Search(2)

	item, err := l.Next(item)

	if err != nil {
		t.Error(err)
	}

	if item.data != 3.0 {
		t.Errorf("Incorrect data. Wanted %v, got %v", 3.0, item.data)
	}

}

func TestSingleDelete(t *testing.T) {
	l := NewLinkedList()

	l.Push("A")
	l.Push("B")
	l.Push("B.5")
	l.Push("C")

	item, _ := l.Search("B.5")

	err := l.Delete(item)

	if err != nil {
		t.Error(err)
	}

	if l.Length != 3 {
		t.Errorf("Incorrect list length, wanted 3, got %v", l.Length)
	}

	item, _ = l.Search("B")

	item, _ = l.Next(item)

	if item.data != "C" {
		t.Errorf("Incorrect data, wanted C, got %v", item.data)
	}
}

func TestSingleDeleteIndex(t *testing.T) {
	l := NewLinkedList()

	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
	l.Push(5)

	err := l.DeleteIndex(4)

	if err != nil {
		t.Error(err)
	}

	if l.Length != 4 {
		t.Errorf("Incorrect length, wanted 4, got %v", l.Length)
	}

	err = l.DeleteIndex(0)
	if err != nil {
		t.Error(err)
	}

	if l.Length != 3 {
		t.Errorf("Incorrect length, wanted 3, got %v", l.Length)
	}

	fmt.Println("Is the search function broken?")

	item, _ := l.Search(2)

	fmt.Println("Nope, not search. Maybe the next function")

	item, err = l.Next(item)

	fmt.Println("This should print if it's not the next function")

	if err != nil {
		t.Error(err)
	}

	if item.data != 3 {
		t.Errorf("Incorrect data, wanted 3, got %v", item.data)
	}

	err = l.DeleteIndex(1)

	if err != nil {
		t.Error(err)
	}

	item, err = l.Search(2)

	if err != nil {
		t.Error(err)
	}

	item, err = l.Next(item)

	if err != nil {
		t.Error(item)
	}

	if item.data != 4 {
		t.Errorf("Incorrect data, wanted 4, got %v", item.data)
	}

}
