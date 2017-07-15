package testing

import (
	"reflect"
	"testing"
)

// in principle expected shouldn't be necessary
// you can use the go standard library sort to test your functions against a known working implementation
type TestCase struct {
	// TODO add name string for easier subtests
	Input, Expected []int
}

type sortFunc func([]int) ([]int, int)
type swapFunc func([]int, []int) []int

// RunSortTest evaluates the functionality of sort function
func (c TestCase) RunSortTest(t *testing.T, sort sortFunc) {
	t.Run("", func(t *testing.T) {
		var output, _ = sort(c.Input)

		if !reflect.DeepEqual(c.Expected, output) {
			t.Errorf("sort(%v) is not equal to %v (should not be actual output %v)", c.Input, c.Expected, output)
		}
	})
}

// RunSwapTests evaluates the functionality of swap function
func (c TestCase) RunSwapTests(t *testing.T, swap swapFunc) {

	t.Run("", func(t *testing.T) {

		var output = swap(c.Input, c.Expected)
		if len(c.Input) != len(output) {
			t.Errorf("Expected array to have length : %v but found to have : %v", len(c.Input), len(output))
		}

		if cap(c.Input) != len(output) {
			t.Errorf("Expected array to have a capacity : %v but found it to have : %v", cap(c.Input), cap(output))
		}

	})

}
