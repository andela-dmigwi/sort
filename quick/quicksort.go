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
		lessThan, greaterThan, equalTo, sorted []int
		elem                                   int
		p                                      = list[0]
	)
	for _, elem = range list {
		if elem > p {
			greaterThan = append(greaterThan, elem)
		} else if elem < p {
			lessThan = append(lessThan, elem)
		} else {
			equalTo = append(equalTo, elem)
		}
	}

	// If only equalTo has values then the list is sorted
	if len(append(lessThan, greaterThan...)) == 0 {
		p = 0
	} else {
		p = len(lessThan, equalTo...))
	}

	sorted = append(append(lessThan, equalTo...), greaterThan...)

	return swapSlice(list, sorted), p
}

func sortRecursively(unsortedArray []int) []int {
	if len(unsortedArray) < 2 {
		return []int{}
	}

	var result, pos = orderByPivot(unsortedArray)

	if len(result[:pos]) > 1 && pos > 0 {
		sortRecursively(result[:pos])
	}

	if len(result[pos:]) > 1 && pos > 0 {
		sortRecursively(result[pos:])
	}

	return result

}

//Sort return an array of integer arrange in ascending order by implementing quickSort algorithm
func Sort(array []int) ([]int, int) {

	sortRecursively(array[:])

	return array, 0
}
