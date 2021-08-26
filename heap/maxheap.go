//package heap

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
func (h *Heap) getParent(i int) int{
	return (i-1)/2
}

//AddEntry adds an entry to the end of the heap
func (h *Heap) Insert(b []byte){
	h.Heap = append(h.Heap, b)
	h.heapifyUp()
}

//Extract removes and returns the max entry in the heap
func (h *Heap) Extract() ([]byte, error){
	if len(h.Heap) < 1{
		return []byte, fmt.Errorf("Unable to extract: heap size is 0")
	}
	res := h.Heap[0]
	h.Heap = h.Heap[0] = h.Heap[len(h.Heap)-1]
	h.Heap = h.Heap[:len(h.Heap)-1]
	h.HeapifyDown()
	return res, nil
}

//heapifyUp sorts the heap after a new entry
func (h *Heap) heapifyUp(){
	
	current := h.Heap[len(h.Heap)-1]

	for h.Heap[h.GetParent(current)] > h.Heap[current]{
		h.Swap(current, next)
		current = next
	}
}

//Swap switches the value of two indices
func (h *Heap) Swap(current, next int){
	h.Heap[current], h.Heap[next] = h.Heap[next], h.Heap[current]
}

//CompareChildren compares the children of a given index and returns the larger. If the right child
//does not exist, the left child is returned
func (h *Heap) CompareChildren(i int) int{
	
	left := h.getLeftChild(i)

	right := h.getRightChild(i)
	
	if left >= len(h.Heap){
		return nil
	}
	if right >= len(h.Heap){
		return left
	}
	if res := bytes.Compare(h.Heap[left], h.Heap[right])
}


//heapifyDown sorts the heap after removing an entry
func (h *Heap) heapifyDown(i int){

	if h.getLeftChild(i) > len(h.Heap){
		return
	}
}