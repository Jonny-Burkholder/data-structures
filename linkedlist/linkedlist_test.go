package linkedlist

import "testing"

func TestSinglePush(t *testing.T) {
	l := NewLinkedList()
	l.Push("one")
	l.Push("two")
	l.Push("three")
	if l.Head.data != "one" {
		t.Errorf("Incorrect item, wanted %s, got %s", "item one", l.Head.data)
	}
	if l.Tail.data != "two" {
		t.Errorf("Incorrect item, wanted %s, got %s", "two", l.Tail.data)
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
	l.Unshift("one")
	l.Unshift("two")
	l.Unshift("three")

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
	l.Unshift(1)
	l.Unshift(2)
	l.Unshift(3)

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
