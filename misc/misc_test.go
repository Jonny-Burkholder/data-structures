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

func testQueueData(t *testing.T) {
	q := NewQueue()
	q.QueueData([]byte("1"))
	if bytes.Compare(q.Queue[0], []byte("1")) != 0 {
		t.Errorf("Queue error: wanted %v, got %v", []byte("1"), q.Queue[0])
	}
	q.QueueData([]byte("2"))
	if bytes.Compare(q.Queue[1], []byte("2")) != 0 {
		t.Errorf("Queue error: wanted %v, got %v", []byte("2"), q.Queue[1])
	}
	if len(q.Queue) != 2 {
		t.Errorf("Queue error: invalid queue length. Wanted %v, got %v", 2, len(q.Queue))
	}
}

func testDequeue(t *testing.T) {
	q := NewQueue()
	if _, err := q.Dequeue(); err.Error() != "Cannot dequeue: queue is empty" {
		t.Error("Dequeue error: did not return empty queue error")
	}
	q.QueueData([]byte("1"), []byte("2"), []byte("3"))
	data, err := q.Dequeue()
	if bytes.Compare(data, []byte("1")) != 0 {
		t.Errorf("Dequeue error: wanted %v, got %v", []byte("1"), data)
	}
	if err != nil {
		t.Errorf("Dequeue error: %s", err.Error())
	}
}
