package unionfind

import "fmt"

//union is for screwing around with, well, union stuff. Mostly union find and all that jazz

//Finder contains a slice of int and implements the quickfind method
type finder struct {
	len int
	arr []int
}

//NewFinder takes an int as an argument, and returns a finder struct such that all the elements
//in the finder array are not joined
func Newfinder(len int) *finder {
	arr := make([]int, len)
	for i := 0; i < len; i++ {
		arr[i] = i
	}
	return &finder{
		len: len,
		arr: arr,
	}
}

//GetValue returns the value at a given index of the finder's array
func (f *finder) GetValue(i int) (int, error) {
	if i < 0 || i > f.len {
		return 0, fmt.Errorf("Error: index %v out of range %v", i, f.len)
	}
	return f.arr[i], nil
}

//Join takes an initial integer argument for an index, then a variadic integer argument for
//indices to be joined to that index. It returns a non-nil error if the index is not in range
func (f *finder) Join(index int, children ...int) error {
	if index < 0 || index > f.len {
		return fmt.Errorf("Error: index %v out of range %v", index, f.len)
	}
	res := f.arr[index] //get value at index
	for _, child := range children {
		if child < 0 || child > f.len { //if child is out of range
			return fmt.Errorf("Error: index %v out of range %v", child, f.len)
		}
		f.arr[child] = res //join indices by changing value of children to value of parent
	}
	return nil
}

//JoinEager takes two integer arguments and joins them together. If the child is part of a group,
//the entire group is joined. Slowly. Because big ol' O is nice and linear here folks, if not worse
//Note that I could have made this variadic, but honestly that's not really going to be any different
//than just calling the function multiple times
func (f *finder) JoinEager(parent, child int) error {
	//if index of parent or child is out of range, return an error
	if parent < 0 || child < 0 || parent > f.len || child > f.len {
		return fmt.Errorf("Error: index out of range %v", f.len)
	}
	//get value of parent group
	res := f.arr[parent]
	//clue tells us what group child is currently in
	clue := f.arr[child]
	for i, v := range f.arr {
		if v == clue { //if the number is in our clue group
			//add it to the parent group
			f.arr[i] = res
		}
	}
	return nil
}

//Connected tells us if there is a connection between two numbers
//Again, I could make this variadic, but see above
func (f *finder) Connected(i, j int) bool { //I realize I'm using super inconsistent naming, but I haven't learned any proper terms here
	return f.arr[i] == f.arr[j]
}
