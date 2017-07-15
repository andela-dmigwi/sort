package quick

func swapSlice(toSort, sorted []int) []int {
	var (
		elem, index int
	)

	for index, elem = range sorted {
		toSort[index] = elem
	}

	return toSort

}

func orderByPivot(list []int) ([]int, int) {
	var (
		lessThan, greaterThan, equalTo []int
		p                              = list[0]
	)
	for _, elem := range list {
		if elem > p {
			greaterThan = append(greaterThan, elem)
		} else if elem < p {
			lessThan = append(lessThan, elem)
		} else {
			equalTo = append(equalTo, elem)
		}
	}
	lessThan = append(lessThan, equalTo...)
	p = len(lessThan)

	lessThan = append(lessThan, greaterThan...)

	return swapSlice(list, lessThan), p
}

// func Sort(unsortedArray []int) []int {

// }
