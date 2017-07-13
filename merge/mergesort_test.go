package merge_test

import (
	"testing"

	"github.com/andela-dmigwi/sort/merge"
	. "github.com/andela-dmigwi/sort/testing"
)

func TestMergeSort(t *testing.T) {
	for _, c := range []TestCase{
		TestCase{[]int{}, []int{}},
		TestCase{[]int{0}, []int{0}},
		TestCase{[]int{0, 2, 1}, []int{0, 1, 2}},
		TestCase{[]int{0, 2, 4, 6, 1, 3, 5}, []int{0, 1, 2, 3, 4, 5, 6}},
		TestCase{[]int{1, 2, 3, 7, 8, 0, 1, 4, 5, 6, 7}, []int{0, 1, 1, 2, 3, 4, 5, 6, 7, 7, 8}},
	} {
		c.Run(t, merge.Sort)

	}
}
