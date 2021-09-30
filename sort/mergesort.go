package sort

//Merge sort splits up arrays and sorts them. Then merges them. You know, like in the name

//Sort takes an input slice of integers and performs the merge sort operation on it
func Sort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	//break array into smaller chunks
	a := arr[:len(arr)-len(arr)/2]
	b := arr[len(a):]
	//sort the chunks
	if len(a) > 1 {
		a = Sort(a)
	}
	if len(b) > 1 {
		b = Sort(b)
	}
	//merge the chunks back together
	return Merge(a, b)
}

//Merge takes smaller, sorted arrays and merges them together
func Merge(a, b []int) []int {
	res := make([]int, len(a)+len(b))
	//if the arrays are one int long, just compare them directly
	if len(a) == 1 && len(b) == 1 {
		if a[0] < b[0] {
			res[0], res[1] = a[0], b[0]
		} else {
			res[0], res[1] = b[0], a[0]
		}
	} else {
		res = make([]int, 0)
		i, j := 0, 0
		for i < len(a) && j < len(b) {
			if a[i] == b[j] {
				res = append(res, a[i])
				i++
				j++
				continue
			}
			if a[i] < b[j] {
				res = append(res, a[i])
				i++
				continue
			}
			if a[i] > b[j] {
				res = append(res, b[j])
				j++
				continue
			}
		}
	}
	return res
}
