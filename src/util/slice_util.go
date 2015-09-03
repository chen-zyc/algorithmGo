package util

func SliceDeleteString(slice []string, index int) []string {
	len := len(slice)
	if index < 0 || index >= len {
		return slice
	}
	if len == 0 {
		return slice
	}
	if len == 1 {
		return slice[0:0]
	}
	if index == 0 {
		return slice[1:]
	}
	if index == len-1 {
		return slice[:len-1]
	}
	return append(slice[:index], slice[index+1:]...)
}

func SliceDeleteInt(slice []int, index int) []int {
	len := len(slice)
	if index < 0 || index >= len {
		return slice
	}
	if len == 0 {
		return slice
	}
	if len == 1 {
		return slice[0:0]
	}
	if index == 0 {
		return slice[1:]
	}
	if index == len-1 {
		return slice[:len-1]
	}
	return append(slice[:index], slice[index+1:]...)
}

func CopySliceInt64(src []int64) []int64 {
	dst := make([]int64, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = src[i]
	}
	return dst
}

func CopySliceInt(src []int) []int {
	dst := make([]int, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = src[i]
	}
	return dst
}
