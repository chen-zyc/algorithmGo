package sorts

// 计数排序，时间复杂度O(max+n)。
// 计数排序是稳定排序。
// in中的元素必须 0 <= in[k] <= max。
func CountingSort(in []uint, max uint) (out []uint) {
	out = make([]uint, len(in))
	c := make([]uint, max+1)
	for i := 0; i < len(in); i++ {
		c[in[i]]++ // 出现一次时是1，从1开始计数的
	}
	// c[i] now contains the number of elements equal to i.
	for i := uint(1); i <= max; i++ {
		c[i] += c[i-1]
	}
	// c[i] now contains the number of elements less than or equal to i.
	for i := len(in) - 1; i >= 0; i-- {
		out[c[in[i]]-1] = in[i] // -1是因为c[j]==1时应该放到[0]中
		c[in[i]]--              // 下一个重复的放到前面去
	}
	return
}

func CountingSort2(in []uint, max uint) (out []uint) {
	out = make([]uint, len(in))
	c := make([]uint, max+1)
	for i := 0; i < len(in); i++ {
		c[in[i]]++ // 出现一次时是1，从1开始计数的
	}
	j := 0
	for i := uint(0); i < uint(len(c)); i++ {
		k := c[i]                      // 元素i出现的次数
		for m := uint(0); m < k; m++ { // 向out中写m个i
			out[j] = i
			j++
		}
	}
	return out
}
