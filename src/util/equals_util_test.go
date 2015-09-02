package util

import (
	"testing"
)

func TestEquals(t *testing.T) {
	s := []string{
		"a", "b", "c",
	}
	stringSlice := StringSlice(s)

	if Equals(stringSlice, StringSlice([]string{})) {
		t.Errorf("%v equals empty slice?", s)
	}
	if Equals(stringSlice, StringSlice([]string{"a"})) {
		t.Errorf("%v equals ['a']?", s)
	}
	if !Equals(stringSlice, StringSlice([]string{"a", "b", "c"})) {
		t.Errorf("%v not equals ['a','b','c']?", s)
	}
	if Equals(stringSlice, StringSlice([]string{"a", "b", "c", "d"})) {
		t.Errorf("%v equals ['a','b','c','d']?", s)
	}
}
