package misc

import "fmt"

//Stack is a slice of byte slices, each byte being a new entry
type Stack struct {
	Stack [][]byte
}

//NewStack creates and returns a new stack with a stack of length 0
func NewStack() *Stack {
	s := make([][]byte, 0)
	return &Stack{
		Stack: s,
	}
}

//Push adds a new item to the stack
func (s *Stack) Push(b ...[]byte) {
	s.Stack = append(s.Stack, b...)
}

//Pop removes and returns the last item in the stack. If the stack is empty, it returns a non-nil error
func (s *Stack) Pop() ([]byte, error) {
	if len(s.Stack) < 1 {
		return []byte{}, fmt.Errorf("Unable to perform operation: stack is empty")
	}
	res := s.Stack[len(s.Stack)]
	s.Stack = s.Stack[len(s.Stack)-1:]
	return res, nil
}
