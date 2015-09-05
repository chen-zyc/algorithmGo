package sort

import (
	"testing"
	"util"
)

func TestMerge(t *testing.T) {
	var tests = []struct {
		arr     []int
		p, q, r int
		res     []int
	}{
		{[]int{1, 2, 3, 5, 6, 7}, 0, 2, 5, []int{1, 2, 3, 5, 6, 7}},
		{[]int{1}, 0, 0, 0, []int{1}},
		{[]int{1, 3, 5, 2, 4}, 0, 2, 4, []int{1, 2, 3, 4, 5}},
		{[]int{5, 2, 3, 4, 1, 0}, 2, 3, 4, []int{5, 2, 1, 3, 4, 0}},
		{[]int{5, 4, 2, 3, 1}, 2, 3, 4, []int{5, 4, 1, 2, 3}},
		{[]int{5, 4, 3, 2, 1}, 1, 1, 2, []int{5, 3, 4, 2, 1}},
	}

	for i, te := range tests {
		merge(MergeIntSlice(te.arr), te.p, te.q, te.r)
		if !util.EqualsIntSlice(te.arr, te.res) {
			t.Errorf("%d: want:%v, but:%v \n", i, te.res, te.arr)
		}
	}
}

func TestMergeSort(t *testing.T) {
	var test = []struct {
		A          []int
		start, end int
		want       []int
	}{
		{[]int{1, 2, 3, 4, 5}, 0, 5, []int{1, 2, 3, 4, 5}},
		{[]int{}, 0, 0, []int{}},
		{[]int{1}, 0, 1, []int{1}},
		{[]int{5, 4, 3, 2, 1}, 1, 4, []int{5, 1, 2, 3, 4}},
	}
	for i, te := range test {
		MergeSortIntSlice(te.A, te.start, te.end+1)
		if !util.EqualsIntSlice(te.A, te.want) {
			t.Errorf("%d: want:%v, but:%v \n", i, te.want, te.A)
		}
	}
}
