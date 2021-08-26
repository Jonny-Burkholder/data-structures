package heap

import (
	"bytes"
	"fmt"
)

//Heap implements the heap abstract data type
type Heap struct {
	Heap [][]byte
}

//NewHeap creates and returns a new heap struct
func NewHeap() *Heap {
	h := make([][]byte, 0)
	return &Heap{
		Heap: h,
	}
}

//getLeftChild takes an integer argument for an index and returns the index of the left child
func (h *Heap) getLeftChild(i int) int {
	return (i * 2) + 1
}

//getRightChild takes an integer argument for an index and returns the index of the right child
func (h *Heap) getRightChild(i int) int {
	return (i * 2) + 2
}

//getParent takes an integer argument for an index and returns the index of the parent node
func (h *Heap) getParent(i int) int {
	return (i - 1) / 2
}

//AddEntry adds an entry to the end of the heap
func (h *Heap) Insert(b []byte) {
	h.Heap = append(h.Heap, b)
	h.heapifyUp()
}

//Extract removes and returns the max entry in the heap
func (h *Heap) Extract() ([]byte, error) {
	if len(h.Heap) < 1 {
		return []byte{}, fmt.Errorf("Unable to extract: heap size is 0")
	}
	res := h.Heap[0]
	h.Heap = h.Heap[:len(h.Heap)-1]
	h.heapifyDown(0)
	return res, nil
}

//heapifyUp sorts the heap after a new entry
func (h *Heap) heapifyUp() {

	current := len(h.Heap) - 1

	parent := h.getParent(current)

	for bytes.Compare(h.Heap[current], h.Heap[parent]) > 0 {
		h.Swap(current, parent)
		current = parent
		parent = h.getParent(current)
	}
}

//Swap switches the value of two indices
func (h *Heap) Swap(current, next int) {
	h.Heap[current], h.Heap[next] = h.Heap[next], h.Heap[current]
}

//heapifyDown sorts the heap after removing an entry
func (h *Heap) heapifyDown(i int) {

	left := h.getLeftChild(i)
	right := h.getRightChild(i)

	//check that there are children. If not, return
	if left > len(h.Heap) {
		return
	} else if right > len(h.Heap) { //if there is a left child, but not right, compare left to parent
		if bytes.Compare(h.Heap[left], h.Heap[i]) > 0 {
			h.Swap(left, i)
		}
	} else {
		//compare children to see which is larger
		child := 0
		if bytes.Compare(h.Heap[left], h.Heap[right]) > 0 {
			child = left
		} else {
			child = right
		}
		//compare larger child to parent
		if bytes.Compare(h.Heap[child], h.Heap[i]) > 0 {
			h.Swap(i, child)
			h.heapifyDown(child)
		}
	}
}
