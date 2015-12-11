package sorts

import (
	"math/rand"
	"testing"
	"time"
	"util"
)

func TestCountingSort(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for j := 0; j < 1; j++ {
		in := make([]uint, 50)
		max := uint(0)
		for i := 1; i < len(in); i++ {
			in[i] = uint(rand.Intn(100))
			if in[i] > max {
				max = in[i]
			}
		}
		out := CountingSort(in, max)
		// check
		if !util.SliceSortedUint(out) {
			t.Errorf("%v is not a sorted array", out)
			break
		}
	}
}

func TestCountingSort2(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for j := 0; j < 1; j++ {
		in := make([]uint, 50)
		max := uint(0)
		for i := 1; i < len(in); i++ {
			in[i] = uint(rand.Intn(100))
			if in[i] > max {
				max = in[i]
			}
		}
		out := CountingSort2(in, max)
		// check
		if !util.SliceSortedUint(out) {
			t.Errorf("%v is not a sorted array", out)
			break
		}
	}
}
