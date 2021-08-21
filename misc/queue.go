package misc

import "fmt"

//Just use a linked list for this

//Queue is a struct with a slice of byte slices
type Queue struct {
	Queue [][]byte
}

//NewQueue returns a new queue
func NewQueue() *Queue {
	return &Queue{
		Queue: make([][]byte, 0),
	}
}

//QueueData adds an item or items to the end of the queue
func (q *Queue) QueueData(b ...[]byte) {
	q.Queue = append(q.Queue, b...)
}

//Dequeue removes an item from the front of the queue and returns it, or returns an non-nil error if
//the queue is empty
func (q *Queue) Dequeue() ([]byte, error) {
	if len(q.Queue) < 1 {
		return []byte{}, fmt.Errorf("Cannot dequeue: queue is empty")
	}
	res := q.Queue[0]
	q.Queue = q.Queue[1:]
	return res, nil
}
