package sort

// 插入排序，时间复杂度：最好最坏都是n^2
//
// 排序范围：[start, end)。
func Insert(a []int, start, end int) {
	if end-start < 2 {
		return
	}
	for front := start + 1; front < end; front++ {
		for back := front; back > start && a[back] < a[back-1]; back-- {
			a[back], a[back-1] = a[back-1], a[back]
		}
	}
}
