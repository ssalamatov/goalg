package qsort

func Partition(arr []int, left, right int) ([]int, int) {
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return arr, i
}

func Sort(arr []int, left, right int) []int {
	if left < right {
		arr, q := Partition(arr, left, right)
		arr = Sort(arr, left, q-1)
		arr = Sort(arr, q+1, right)
		return arr
	}
	return arr
}

func SortArr(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	return Sort(arr, 0, len(arr)-1)
}
