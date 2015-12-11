package sorts

import (
	"reflect"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var tests = []struct {
		in, want []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 3, 2}, []int{1, 2, 3}},
		{[]int{-5, 3, 100, 43, 41}, []int{-5, 3, 41, 43, 100}},
		{[]int{20, 14, 2, 7, 1, 0}, []int{0, 1, 2, 7, 14, 20}},
	}
	for i, test := range tests {
		QuickSort(sort.IntSlice(test.in), nil)
		if !reflect.DeepEqual(test.in, test.want) {
			t.Errorf("#%d: want %v, actual: %v", i, test.want, test.in)
		}
	}
}
