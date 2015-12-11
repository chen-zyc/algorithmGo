package sorts

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
	"util"
)

var bigData []int
var sortedData []int

func init() {
	rand.Seed(time.Now().Unix())
	var dataLen int = 1e5
	bigData = make([]int, dataLen)
	for i := 0; i < len(bigData); i++ {
		bigData[i] = rand.Int()
	}
}
func geneBigData() []int {
	return util.CopySliceInt(bigData)
}

func checkSortedData(sorted []int) {
	if sortedData == nil {
		sortedData = sorted
	} else {
		if !reflect.DeepEqual(sortedData, sorted) {
			panic("soted data not equals.")
		}
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bigData := geneBigData()
		InsertionSort(sort.IntSlice(bigData), 0, len(bigData))
        checkSortedData(bigData)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bigData := geneBigData()
		HeapSort(sort.IntSlice(bigData), 0, len(bigData))
        checkSortedData(bigData)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bigData := geneBigData()
		MergeSortIntSlice(bigData, 0, len(bigData))
        checkSortedData(bigData)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bigData := geneBigData()
		QuickSort(sort.IntSlice(bigData), nil)
        checkSortedData(bigData)
	}
}
