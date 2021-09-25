package misc

//The goal of list-sorted is to make a function that takes a variadic number of input arrays and
//prints the elements in the arrays in sorted order

//List takes inputs arrs and prints them in sorted order. If there are duplicate entries, the entry
//from the latest input is used
func List(arrs ...[]int) []int {
	res := []int{}
	if len(arrs) < 1 {
		//go to end of if statement
	} else if len(arrs) == 1 {
		res = arrs[0]
	} else {
		res = arrs[0] //I know this is duplicate code, but I need this to be empty if arrs is also empty
		for _, arr := range arrs {
			//range through the slice and see if each number is contained in slice res
			for _, v := range arr {
				//if the number is already in slice res
				if index, ok := isIn(v, res); ok != false {
					//replace the number in res with the number from the current slice
					res[index] = v
				} else {
					res = append(res, v)
				}
			}
		}
	}

	//now let's sort res
	for i := range res {
		for j := range res {
			if res[i] < res[j] {
				res[i], res[j] = res[j], res[i]
			}
		}
	}

	return res
}

//isIn takes an integer input and an array input. If the integer is found in the array, the index
//of the number within the array is returned, along with true. If the integer is not found in the
//array, 0 and false are returned
func isIn(a int, arr []int) (int, bool) {
	for i, v := range arr {
		if v == a {
			return i, true
		}
	}
	return 0, false
}
