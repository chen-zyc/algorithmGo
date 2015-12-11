package sorts

import (
	"sort"
	"testing"
	"util"
)

func TestInsertionSort(t *testing.T) {
	var tests = []struct {
		in, want   []int
		start, end int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, 0, 3},
		{[]int{}, []int{}, 0, 0},
		{[]int{1}, []int{1}, 0, 1},
		{[]int{1, 3, 2}, []int{1, 2, 3}, 0, 3},
		{[]int{-5, 3, 100, 43, 41}, []int{-5, 3, 100, 43, 41}, 0, 3},
		{[]int{-5, 3, 100, 43, 41}, []int{-5, 3, 41, 43, 100}, 0, 5},
		{[]int{20, 14, 2, 7, 1, 0}, []int{20, 14, 0, 1, 2, 7}, 2, 6},
		{[]int{20, 14, 2, 7, 1, 0}, []int{20, 2, 7, 14, 1, 0}, 1, 4},
		{[]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}, []int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}, 0, 10},
		{[]int{1, 7, 3, 5, 0, 2, 8, 4}, []int{0, 1, 2, 3, 4, 5, 7, 8}, 0, 8},
	}

	for i, test := range tests {
		InsertionSort(sort.IntSlice(test.in), test.start, test.end)

		if !util.Equals(util.IntSlice(test.in), util.IntSlice(test.want)) {
			t.Errorf("%d: want: %v, actually: %v \n", i, test.want, test.in)
		}
	}
}

func TestHeapSort(t *testing.T) {
	var tests = []struct {
		in, want   []int
		start, end int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, 0, 3},
		{[]int{}, []int{}, 0, 0},
		{[]int{1}, []int{1}, 0, 1},
		{[]int{1, 3, 2}, []int{1, 2, 3}, 0, 3},
		{[]int{-5, 3, 100, 43, 41}, []int{-5, 3, 100, 43, 41}, 0, 3},
		{[]int{-5, 3, 100, 43, 41}, []int{-5, 3, 41, 43, 100}, 0, 5},
		{[]int{20, 14, 2, 7, 1, 0}, []int{20, 14, 0, 1, 2, 7}, 2, 6},
		{[]int{20, 14, 2, 7, 1, 0}, []int{20, 2, 7, 14, 1, 0}, 1, 4},
		{[]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}, []int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}, 0, 10},
		{[]int{1, 7, 3, 5, 0, 2, 8, 4}, []int{0, 1, 2, 3, 4, 5, 7, 8}, 0, 8},
	}

	for i, te := range tests {
		HeapSort(sort.IntSlice(te.in), te.start, te.end)
		if !util.Equals(util.IntSlice(te.in), util.IntSlice(te.want)) {
			t.Errorf("$d: want: %v, but %v \n", i, te.want, te.in)
		}
	}
}
