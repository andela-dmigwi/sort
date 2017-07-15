package quick

import (
	"testing"

	. "github.com/andela-dmigwi/sort/testing"
)

func TestSwapSlice(t *testing.T) {
	var elem TestCase

	for _, elem = range []TestCase{
		TestCase{[]int{0}, []int{0}},
		TestCase{[]int{2, 1}, []int{1, 2}},
		TestCase{[]int{1, 0, -1}, []int{-1, 0, 1}},
		TestCase{[]int{6, 7, 9, 8, 9}, []int{6, 7, 8, 9, 9}},
		TestCase{[]int{6, 2, 5, 9, 2, 1, 0, 6}, []int{0, 1, 2, 2, 5, 6, 1, 9}},
		TestCase{[]int{6, 2, 5, 9, 2, 1, 0, 6}, []int{0, 1, 2, 2, 5, 6, 1, 9}},
		TestCase{[]int{6, 2, 5, 9}, []int{0, 1, 2, 2, 5, 6, 1, 9}[:3]},
		TestCase{[]int{6, 2, 5, 9, 2}, []int{0, 1, 2, 2, 5, 6, 1, 9}[:4]},
	} {

		elem.RunSwapTests(t, swapSlice)

	}

}

func TestOrder(t *testing.T) {
	var elem TestCase

	for _, elem = range []TestCase{
		TestCase{[]int{1, 0, -1}, []int{0, -1, 1}},
		TestCase{[]int{6, 7, 9, 8, 9}, []int{6, 7, 9, 8, 9}},
		TestCase{[]int{6, 2, 5, 9, 2, 1, 0, 6}, []int{2, 5, 2, 1, 0, 6, 6, 9}},
		TestCase{[]int{6, 2, 5, 9, 2}, []int{2, 5, 2, 6, 9}},
	} {
		elem.RunSortTest(t, orderByPivot)
	}
}

func TestSort(t *testing.T) {
	var elem TestCase

	for _, elem = range []TestCase{
		TestCase{[]int{1, 0, -1}, []int{-1, 0, 1}},
		TestCase{[]int{3, 1, 0, 6, -1}, []int{-1, 0, 1, 3, 6}},
		TestCase{[]int{6, 7, 9, 8, 9}, []int{6, 7, 8, 9, 9}},
		TestCase{[]int{6, 2, 5, 9, 2, 1, 0, 6}, []int{0, 1, 2, 2, 5, 6, 6, 9}},
		TestCase{[]int{6, 2, 5, 9, 2}, []int{2, 2, 5, 6, 9}},
		TestCase{[]int{6, 2, 5, 9, 2}, []int{2, 2, 5, 6, 9}},
		TestCase{[]int{6, 6, 6, 9, 2}, []int{2, 6, 6, 6, 9}},
	} {
		elem.RunSortTest(t, Sort)
	}

}
