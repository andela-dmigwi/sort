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

type sortFunc func([]int) []int

func (c TestCase) Run(t *testing.T, sort sortFunc) {
	t.Run("", func(t *testing.T) {
		output := sort(c.Input)

		if !reflect.DeepEqual(c.Expected, output) {
			t.Errorf("sort(%v) != %v (should not be actual output %v)", c.Input, c.Expected, output)
		}
	})
}
