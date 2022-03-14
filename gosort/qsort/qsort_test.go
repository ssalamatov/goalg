package qsort

import (
	"testing"

	"github.com/ssalamatov/goalg/gosort/internal/comparable"
)

func TestQSort(t *testing.T) {
	arr := []int{4, 1, 4, 78, 9, 1, 2}
	sortedArr := []int{1, 1, 2, 4, 4, 9, 78}

	if !comparable.IsEqualSlice(SortArr(arr), sortedArr) {
		t.Errorf("Not equal %v %v", SortArr(arr), sortedArr)
	}
}
