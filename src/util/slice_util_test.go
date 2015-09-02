package util

import "testing"

func TestSliceDeleteString(t *testing.T) {
	slice := []string{"a", "b", "c"}

	slice2 := SliceDeleteString(slice, -1)
	if !Equals(StringSlice(slice), StringSlice(slice2)) {
		t.Errorf("delete %v at %d err: %v", slice, -1, slice2)
	}

	slice2 = SliceDeleteString(slice, 3)
	if !Equals(StringSlice(slice), StringSlice(slice2)) {
		t.Errorf("delete %v at %d err: %v", slice, 3, slice2)
	}

	slice2 = SliceDeleteString(slice, 0)
	if !Equals(StringSlice(slice2), StringSlice([]string{"b", "c"})) {
		t.Errorf("delete at %d err: %v", 0, slice2)
	}

	slice = []string{"a", "b", "c"}
	slice2 = SliceDeleteString(slice, 1)
	if !Equals(StringSlice(slice2), StringSlice([]string{"a", "c"})) {
		t.Errorf("delete at %d err: %v", 1, slice2)
	}

	slice = []string{"a", "b", "c"}
	slice2 = SliceDeleteString(slice, 2)
	if !Equals(StringSlice(slice2), StringSlice([]string{"a", "b"})) {
		t.Errorf("delete at %d err: %v", 2, slice2)
	}
}
