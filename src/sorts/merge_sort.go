package sorts

/**************** 归并排序 ******************/

type MergeIntf interface {
	// 容器长度
	Len() int
	// 创建一个length长的容器
	New(length int) MergeIntf
	// this[i] <= other[j]
	LessEqual(i int, other MergeIntf, j int) bool
	// 将other[j]赋值给this[i], this[i] = other[j]
	Set(i int, other MergeIntf, j int)
}

// 归并排序，A[start, end)，时间复杂度:最好最坏都是nlgn+n
func MergeSort(A MergeIntf, start, end int) {
	if 0 <= start && end-start > 1 && end <= A.Len() {
		q := (start + end) / 2 // 分界点
		MergeSort(A, start, q)
		MergeSort(A, q, end)
		merge(A, start, q-1, end-1)
	}
}

// 将A[p, q]和A[q+1, r]合并，这里假设A[p, q]和A[q+1, r]都是已经排好序的
func merge(A MergeIntf, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q

	// 保证两个数组的长度都大于0，没有做数组下标越界检查
	if n1 < 1 && n2 < 1 {
		return
	}

	L := A.New(n1)
	R := A.New(n2)

	for i := 0; i < n1; i++ {
		//		L[i] = A[p+i]
		L.Set(i, A, p+i)
	}

	for i := 0; i < n2; i++ {
		//		R[i] = A[q+1+i]
		R.Set(i, A, q+1+i)
	}

	i, j, k := 0, 0, p
	for ; k <= r; k++ {
		// 有一个数组遍历完了就跳出
		// 这样下面就不需要再判断是不是下标越界了
		if i >= n1 || j >= n2 {
			break
		}
		if L.LessEqual(i, R, j) { // L[i] <= R[j]
			//			A[k] = L[i]
			A.Set(k, L, i)
			i++
		} else {
			//			A[k] = R[j]
			A.Set(k, R, j)
			j++
		}
	}

	// 此时，至多有一个数组未遍历完
	if i < n1 {
		for ; i < n1; i++ {
			A.Set(k, L, i)
			k++
		}
	} else if j < n2 {
		for ; j < n2; j++ {
			A.Set(k, R, j)
			k++
		}
	}
}

/************* 实现 MergeIntf 接口 **************/

type MergeIntSlice []int

func (this MergeIntSlice) Len() int { return len(this) }

func (this MergeIntSlice) New(length int) MergeIntf { return MergeIntSlice(make([]int, length)) }

func (this MergeIntSlice) LessEqual(i int, other MergeIntf, j int) bool {
	if o, ok := other.(MergeIntSlice); ok {
		return this[i] <= o[j]
	}
	panic("must be MergeIntSlice")
}

func (this MergeIntSlice) Set(i int, other MergeIntf, j int) {
	if o, ok := other.(MergeIntSlice); ok {
		this[i] = o[j]
		return
	}
	panic("must be MergeIntSlice")
}

func MergeSortIntSlice(a []int, start, end int) {
	MergeSort(MergeIntSlice(a), start, end)
}
