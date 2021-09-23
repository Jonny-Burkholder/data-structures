package sort

import "testing"

func TestSort(t *testing.T) {

	arr := []int{5, 1, 8, 2, 7}
	sorted := []int{1, 2, 5, 7, 8}
	arr = Sort(arr)
	for i, v := range arr {

		if v != sorted[i] {
			t.Errorf("Array not sorted. Wanted %v, got %v", sorted, arr)
			break
		}
	}
}
