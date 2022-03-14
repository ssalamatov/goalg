package merge

func Merge(left []int, right []int) (buf []int) {
	leftln := len(left)
	rightln := len(right)

	i, j := 0, 0
	for i < leftln && j < rightln {
		if left[i] < right[j] {
			buf = append(buf, left[i])
			i++
		} else {
			buf = append(buf, right[j])
			j++
		}
	}
	for ; i < leftln; i++ {
		buf = append(buf, left[i])
	}
	for ; j < rightln; j++ {
		buf = append(buf, right[j])
	}
	return
}

func SortArr(arr []int) []int {
	ln := len(arr)
	if ln < 2 {
		return arr
	}
	middle := ln / 2
	left := SortArr(arr[:middle])
	right := SortArr(arr[middle:])
	return Merge(left, right)
}
