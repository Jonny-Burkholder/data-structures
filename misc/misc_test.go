package misc

import (
	"bytes"
	"testing"
)

func TestPush(t *testing.T) {
	s := NewStack()
	s.Push([]byte("try"))
	if bytes.Compare(s.Stack[0], []byte("try")) != 0 {
		t.Errorf("Push Error: Wanted %v, got %v", []byte("try"), s.Stack[0])
	}
	s.Push([]byte("try_2"))
	if bytes.Compare(s.Stack[1], []byte("try_2")) != 0 {
		t.Errorf("Push Error: Wanted %v, got %v", []byte("try2"), s.Stack[1])
	}
	s.Push([]byte("3"), []byte("4"))
	if len(s.Stack) != 4 {
		t.Errorf("Push error: invalid stack length. Wanted %v, got %v", 4, len(s.Stack))
	}
}

func testPop(t *testing.T) {
	s := NewStack()
	_, err := s.Pop()
	if err.Error() != "Unable to perform operation: stack is empty" {
		t.Error("Pop error: did not return empty stack error")
	}
	s.Push([]byte("One data"), []byte("Two data"), []byte("Three data"), []byte("Four data, ah, ah, ah"))
	data, err := s.Pop()
	if err != nil {
		t.Error(err.Error())
	}
	if bytes.Compare(data, []byte("Four data, ah, ah, ah")) != 0 {
		t.Errorf("Pop error: wanted %v, got %v", []byte("Four data, ah, ah, ah"), data)
	}
	if len(s.Stack) != 3 {
		t.Errorf("Pop error: invalid stack length. Wanted %v, got %v", 3, len(s.Stack))
	}
}
