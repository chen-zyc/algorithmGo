package sort

import (
	"math/rand"
	"sort"
	"time"
)

// QuickSort 快速排序，范围是A[p, r]。最坏情况下O(n^2)，与插入排序相同，最好情况下O(nlgn)，与归并排序相同。
// 最坏情况是partition划分的两个子数组中有一个是0长度的，比如数组本来就是已经排好序的了，而这种情况下插入排序的时间复杂度是O(n)。
// 最好情况是partition划分的两个数组的规模都不大于n/2，此时时间复杂度是O(nlgn)。
// 即使划分的两个子数组的规模是99:1，时间复杂度依旧是O(nlgn)，所以，大部分情况下快速排序的复杂度都是O(nlgn)。
func QuickSort(data sort.Interface, pivotSelector PivotSelector) {
	if pivotSelector == nil {
		pivotSelector = defaultSelectPivot
	}
	doQuickSort(data, 0, data.Len()-1, pivotSelector)
}

// PivotSelector 在data[start, end]范围内选择一个主元。
type PivotSelector func(data sort.Interface, start, end int) int

var defaultSelectPivot = func(data sort.Interface, start, end int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(end-start) + start
}

// 在 [start, end] 范围内排序
func doQuickSort(data sort.Interface, start, end int, pivotSelector PivotSelector) {
	if start < end { // 确保至少有两个元素
		mid := partition(data, start, end, pivotSelector)
		// 除了中间的元素，剩下的分为两组进行排序
		doQuickSort(data, start, mid-1)
		doQuickSort(data, mid+1, end)
	}
}

// 划分data[p,r]，找到一个元素（主元），它前面的都比它小，它后面的都比它大,返回这个元素的位置。
// 时间复杂度是O(n)，其中n是(r-p+1)。
func partition(data sort.Interface, start, end int, pivotSelector PivotSelector) int {
	pivot := pivotSelector(data, start, end) // 主元位置
	data.Swap(pivot, end)                    // 把主元放到最后
	pivot = end
	little := start - 1            // 比主元小的元素最靠近主元那侧的位置
	for i := start; i < end; i++ { // 遍历主元之前的元素
		if data.Less(i, pivot) { // data[i] < 主元
			// 把data[i]放到little队列最后
			little++
			data.Swap(little, i)
		}
	}
	data.Swap(little+1, pivot) // 把主元放到中间
    return little+1 // 返回主元的位置
}
