package sort

import (
	"sort"
)

/**************** 插入排序 ******************/

// InsertionSort 插入排序，时间复杂度O(n^2)。
//
// 这个实现是拷贝的sort包下的 insertionSort()。
func InsertionSort(data sort.Interface, start, end int) {
	for i := start + 1; i < end; i++ {
		for j := i; j > start && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

/**************** 堆排序 ******************/

// HeapSort 堆排序，时间复杂度O(nlgn)。
//
// 拷贝自sort包下的 heapSort()。
func HeapSort(data sort.Interface, a, b int) {
	// first是第一个数的偏移量，下面的计算是从0开始的，操作数据的时候要加上偏移量
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	// (hi - 1) / 2 是 hi 的父节点
	// i 先从最后一个非叶子节点开始，但最后一个非叶子节点可以用 (hi - 1 - 1) / 2，不明白为什么这样？
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		// 把最大的放到最后
		data.Swap(first, first+i)
		// 除最后以排序的，重新调整堆使其满足最大堆性质
		siftDown(data, lo, i, first)
	}
}

// siftDown implements the heap property on data[lo, hi).
// first is an offset into the array where the root of the heap lies.
// 确保以 data[lo] 为根节点的子树满足最大堆的性质。
func siftDown(data sort.Interface, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1 // 左子节点
		if child >= hi {
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}
