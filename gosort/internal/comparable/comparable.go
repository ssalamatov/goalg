package comparable

func IsEqualSlice(x []int, y []int) bool {
	if len(x) != len(y) {
		return false
	}

	for idx, xv := range x {
		if xv != y[idx] {
			return false
		}
	}
	return true
}
