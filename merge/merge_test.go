package merge

import (
	"reflect"
	"testing"
)

type testCase struct {
	a, b, expected []int
}

func TestMerge(t *testing.T) {
	for _, c := range []testCase{
		testCase{[]int{}, []int{}, []int{}},
		testCase{[]int{0}, []int{}, []int{0}},
		testCase{[]int{}, []int{0}, []int{0}},
		testCase{[]int{0}, []int{0}, []int{0, 0}},
		testCase{[]int{1}, []int{0}, []int{0, 1}},
		testCase{[]int{0}, []int{1}, []int{0, 1}},
		testCase{[]int{0, 2}, []int{1}, []int{0, 1, 2}},
		testCase{[]int{0, 2, 4, 6}, []int{1, 3, 5}, []int{0, 1, 2, 3, 4, 5, 6}},
		testCase{[]int{1, 2, 3, 7, 8}, []int{0, 1, 4, 5, 6, 7}, []int{0, 1, 1, 2, 3, 4, 5, 6, 7, 7, 8}},
	} {
		t.Run("", func(t *testing.T) {
			output := slurp(merge(makeChan(c.a), makeChan(c.b)))

			if !reflect.DeepEqual(c.expected, output) {
				t.Errorf("Merge(%v, %v) != %v (should not be actual output %v)", c.a, c.b, c.expected, output)
			}
		})
	}
}
