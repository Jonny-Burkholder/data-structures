package unionfind

import "fmt"

//quickfind implements the quick find search method for a union array

//quickfinder implements the quick find structure and methods. It is otherwise identical
//to the finder struct
type quickfinder struct {
	len int
	arr []int
}

//Newquickfinder takes an integer argument len and returns a new quickfinder struct
//of lenght len, such that all elements in the structs array point to themselves
func Newquickfinder(len int) *quickfinder {
	arr := make([]int, len)
	for i := 0; i < len; i++ {
		arr[i] = i
	}
	return &quickfinder{
		len: len,
		arr: arr,
	}
}

//GetValue takes an integer value index and returns the value at that index
//The function returns a non-nill error if the index is out of range
func (q *quickfinder) GetValue(index int) (int, error) {
	if !q.inrange(index) {
		return 0, fmt.Errorf("Error: index %v out of range %v", index, q.len)
	}
	return q.arr[index], nil
}

//inrange checks to see if a set of indices are in range of the quickfinder's array
//I could have it return an index for error catching, but that wouldn't be as syntactically sexy
func (q *quickfinder) inrange(nums ...int) bool {
	for _, num := range nums {
		if num < 0 || num > q.len {
			return false
		}
	}
	return true
}

//QuickJoin takes two integer arguments as indices, and joins them such that the second index points
//to the first. If either index is out of range, function will return a non-nil error
func (q *quickfinder) QuickJoin(a, b int) error {
	if !q.inrange(a, b) { //if index is out of range
		return fmt.Errorf("Error: index out of range %v", q.len)
	}
	q.arr[b] = a //value of b will now point to the index of a
	return nil
}

//QuickJoinEager does the same thing as QuickJoin, but eagerly
func (q *quickfinder) QuickJoinEager(a, b int) error {
	if !q.inrange(a, b) {
		return fmt.Errorf("Error: index out of range %v", q.len)
	}
	if q.arr[b] == b {
		q.arr[b] = a
		return nil
	}
	return q.QuickJoinEager(a, q.arr[b])
}

//Connected takes two integer arguments and returns a bool of whether or not those indices are
//connected. If one or both indices is out of range, function will return false
func (q *quickfinder) Connected(a, b int) bool {
	if !q.inrange(a, b) {
		return false
	}
	//let's get to the root of a
	for q.arr[a] != a { //for as long as we're not at the root
		a = q.arr[a]
	}
	//rinse and repeat with b
	for q.arr[b] != b { //for as long as we're not at the root
		b = q.arr[b]
	}
	return a == b //if they have the same root, they're connected
}
