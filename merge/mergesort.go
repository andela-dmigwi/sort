package merge

// Sort will return a slice of integers arranged in ascending order based on the input slice
// This is a mergesort that implement using channels concurrentlyg
func Sort(input []int) []int {

	return slurp(mergeSortChan(input))
}

func mergeSortChan(input []int) <-chan int {
	if len(input) <= 1 {
		return makeChan(input)
	}

	var a, b = input[0 : len(input)/2], input[len(input)/2:]

	var sa, sb = mergeSortChan(a), mergeSortChan(b)

	return merge(sa, sb)
}

func merge(a, b <-chan int) <-chan int {
	// TODO clean up code and make sure it's easy to read since this is just the frst iteration
	var (
		a1, b1       int
		haveA, haveB bool

		c = make(chan int)
	)

	go func() {
		defer close(c)

		for {
			if !haveA {
				a1, haveA = <-a
			}
			if !haveB {
				b1, haveB = <-b
			}

			if !haveA && !haveB {
				break
			}

			if haveA && (haveB && (a1 < b1) || !haveB) {
				c <- a1
				haveA = false
			} else { // haveB must be true due to above conditional,
				c <- b1
				haveB = false
			}

		}
	}()

	return c
}

// TODO clean up helper functions for converting arrays to channels
func makeChan(a []int) <-chan int {
	var (
		v int
		c = make(chan int)
	)

	go func() {
		defer close(c)
		for _, v = range a {
			c <- v
		}
	}()

	return c
}

func slurp(c <-chan int) []int {
	var (
		v int

		ret = []int{}
	)

	for v = range c {
		ret = append(ret, v)
	}

	return ret
}
