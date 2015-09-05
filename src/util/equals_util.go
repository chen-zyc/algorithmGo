package util

type EqualsInterface interface {
	Len() int
	Equals(intf EqualsInterface, i int) bool
}

func Equals(i1, i2 EqualsInterface) bool {
	len1, len2 := i1.Len(), i2.Len()
	// 1. 长度不一样
	if len1 != len2 {
		return false
	}
	// 2. 空的
	if len1 == 0 {
		return true
	}
	// 3. 挨个比较
	for i := 0; i < len1; i++ {
		if !i1.Equals(i2, i) {
			return false
		}
	}
	return true
}

func EqualsIntSlice(s1, s2 []int) bool {
	return Equals(IntSlice(s1), IntSlice(s2))
}

/*********** 包装的一些常用切片 ****************/

type StringSlice []string

func (this StringSlice) Len() int {
	return len(this)
}

func (this StringSlice) Equals(intf EqualsInterface, i int) bool {
	if other, ok := intf.(StringSlice); ok {
		return this[i] == other[i]
	} else {
		return false
	}
}

type IntSlice []int

func (this IntSlice) Len() int {
	return len(this)
}
func (this IntSlice) Equals(intf EqualsInterface, i int) bool {
	if other, ok := intf.(IntSlice); ok {
		return this[i] == other[i]
	} else {
		return false
	}
}
