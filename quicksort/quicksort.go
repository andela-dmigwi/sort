package main

import "fmt"

var c = make(chan []int)

func recursive(list *[]int) {
	var (
		lessThan, greaterThan []int
		v                     = *list
		pivot                 = v[0]
	)
	go func() {
		for i := range v {
			if v[i] > pivot {
				greaterThan = append(greaterThan, v[i])
			} else if v[i] < pivot {
				lessThan = append(lessThan, v[i])
			}
		}
		lessThan = append(lessThan, pivot)
		fmt.Println("Less Than ", lessThan, " Greater than ", greaterThan)
		*list = append(lessThan, greaterThan...)
		v = *list

		fmt.Println(" all <<<<<<< Cap:", cap(v), "     ", v)
		c <- v[:len(lessThan)]
		c <- v[len(lessThan):]

	}()

}

func cutSlice(mainArray, subArray []int) []int {
	fmt.Println("Main Array : ", mainArray, " SubArray : ", subArray)
	var pos []int
	for i := range mainArray {
		switch {
		case mainArray[i] == subArray[0]:
			pos = append(pos, i)

		case mainArray[i] == subArray[len(subArray)-1]:
			pos = append(pos, i)
		}
	}
	fmt.Println("Pos  >>", pos)
	return pos
}

func main() {
	var (
		array         = []int{6, 4, 8, 3, 5, 9, 2, 7, 1}
		ora1, ora2, f []int
		// l                     int
	)
	f = array
	recursive(&array)
	// list := &array
	ora1 = <-c
	ora2 = <-c

	// fmt.Println(" all <<<<<<< ", cap(array), "     >>>>>>", array)

	for i := 0; i < len(f); i++ {

		// l = 0
		// *list = append(ora1, ora2...)
		fmt.Println(array)
		// *list = append(a1, a2...)

		// v := *list
		// fmt.Printf(" >>> %v\n\n", v)
		// ora1 = v[:l]
		// ora2 = v[l:]

		if len(ora1) > 1 {
			pos := cutSlice(array, ora1)
			ora1 := array[pos[0] : pos[1]+1]

			fmt.Printf("Append 1 %v Len : %v cap : %v\n", ora1, len(ora1), cap(ora1))
			recursive(&ora1)
			fmt.Printf(" >>> %v\n\n", append(ora1, ora2...))
		}

		if len(ora2) > 1 {
			pos := cutSlice(array, ora2)
			ora2 := array[pos[0] : pos[1]+1]

			fmt.Printf("Append 2 %v Len : %v cap : %v\n", ora2, len(ora2), cap(ora2))
			recursive(&ora2)
			fmt.Printf(" >>> %v\n\n", append(ora1, ora2...))
		}

		ora1 = <-c
		ora2 = <-c

	}

	fmt.Println("Sorted Array ", array)

}
