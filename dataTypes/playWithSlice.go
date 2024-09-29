package dataTypes

import "fmt"

//________________________________________________________________
// 				GO SLICES CONCEPTS
//________________________________________________________________

// Slices represent variable-length sequences whose elements all have the same type.
// 	A slice type is written []T, where the elements have type T;
// 		it looks like an array type without a size.

// A slice is a lightweight data structure that gives access to a subsequence
// (or perhaps all) of the elements of an array, which is known as the
// slice’s underlying array.

// 	A slice has three components: a pointer, a length, and a capacity.
// 		The pointer points to the first element of the array that is
//          reachable through the slice, which is not necessarily the array’s
//          first element.
// 		The length is the number of slice elements; it can’t exceed the capacity,
// 			which is usually the number of elements between
// 			the start of the slice and the end of the underlying array.
// 		The built-in functions len and cap return those values.

// The slice operator s[ i : j ], where 0 ≤ i ≤ j ≤ cap(s),
// 	Creates a new slice that refers to elements i through j-1 of the sequence s,
// 	which may be an array variable, a pointer to an array, or anotherslice.
// 	The resulting slice has j-i elements.
// 	If i is omitted,it’s 0,and if j isomitted, it’s len(s).

// 	Thus the slice months[1:13] refers to the whole range of valid months,
// 	as does the slice months[1:]; the slice months[:] refers to the whole array.

func doSomething(some []int) {
	for i, j := 0, len(some)-1; i < j; i, j = i+1, j-1 {
		some[i], some[j] = some[j], some[i]
	}
}

func playWithSlice() {
	// Way1: Slice declaration & initialization
	slice := [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 99}

	some1 := slice[0:4]
	fmt.Println("Slice: ", slice)
	doSomething(some1)
	fmt.Println("some1: ", slice)

	some2 := slice[4:9]
	fmt.Println("Slice: ", slice)
	doSomething(some2)
	fmt.Println("some2: ", slice)

	some3 := slice[:]
	fmt.Println("Slice: ", slice)
	doSomething(some3)
	fmt.Println("some3: ", slice)

	// Way2: Slice declaration & initialization
	slice2 := make([]int, 3, 5) // By using make() function we can creates a slice2 of length 3 and capacity 5: [0 0 0]
	// The built-in function make() creates a slice of a specified element type,length, and capacity.
	// The capacity argument may be omitted, in which case the capacity equals the length.
	// 		make([]T, len)
	// 		make([]T, len, cap) // same as make([]T, cap)[:len]
	slice2[0] = 1
	slice2[1] = 11
	slice2[2] = 111
	fmt.Println("Value of Slice2: ", slice2)
	fmt.Println("Length of slice2: ", len(slice2))
	fmt.Println("Capacity of slice2: ", cap(slice2))

	slice2 = append(slice2, 4, 5) //  Adds new elements to the slice. If the slice's capacity is not enough.
	// When length & capasity becomes same, then automaically capacity besomes doubled
	fmt.Println("slice after appending: ", slice2)

	copy(slice[0:5], slice2) // deep copy
	fmt.Println("Copy Arr: ", slice)
}

func printCommon(rainy []string, quarter []string) {
	for _, rMonth := range rainy {
		for _, quarterMonth := range quarter {
			if rMonth == quarterMonth {
				println("Month Appear In Both: ", rMonth) // With println we can't apply formatters.
			}
		}
	}
}

// Following Both Signatures Are Equivalent
// func slicesEqual( x []string, y []string ) bool {}
func slicesEqual(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for index := range x {
		if x[index] != y[index] {
			return false
		}
	}
	return true
}

func playWithSliceComparision() {
	months := [...]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	quarter2 := months[4:7]
	rainy := months[6:9]

	fmt.Println("Quarter2: ", quarter2)
	fmt.Println("Rainy: ", rainy)
	printCommon(rainy, quarter2)

	fmt.Println("Months: ", months)
	firstThree := months[1:4]
	fmt.Println("firstThree: ", firstThree)
	quater1 := months[1:4]
	fmt.Println("quater1: ", quater1)

	// invalid operation: firstThree == quater1 (slice can only be compared to nil)
	// fmt.Println("Slices Are Equal: ", firstThree == quater1 )
	result := slicesEqual(firstThree, quater1)
	fmt.Println("Slices Are Equal: ", result)
}

func playWithSliceCap() {
	s := make([]string, 3)
	fmt.Printf("Slice Type: %T \n", s)

	fmt.Println("Slice of string: ", s)
	fmt.Println("Slice Length and Capacity: ", len(s), cap(s))

	s[0] = "Ding"
	s[1] = "Dong"
	s[2] = "Ting"

	fmt.Println("Slice: ", s)
	fmt.Println("Slice Length and Capacity: ", len(s), cap(s))

	s = append(s, "Tong")
	fmt.Println("Slice: ", s)
	fmt.Println("Slice Length and Capacity: ", len(s), cap(s))
	s = append(s, "Ming", "Mong")
	fmt.Println("Slice: ", s)
	fmt.Println("Slice Length and Capacity: ", len(s), cap(s))
	s = append(s, "Modi")

	fmt.Println("Slice: ", s)
	fmt.Println("Slice Length and Capacity: ", len(s), cap(s))
}

func appendInt(x []int, y int) []int {
	var z []int

	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else { // Underlining Array Is Full
		// When Array Is Full, Allocate A New Array
		//	To Append Additional Data
		// fmt.Println("Underlining Array Is Full")
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	z[len(x)] = y
	return z
}

func playWithAppendInt() {
	var x, y []int

	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf(" %d Capacity=%d \t %v \n", i, cap(y), y)
		x = y
	}
}

// OR

func appendSlice(x []int, y ...int) []int {
	var z []int

	zlen := len(x) + len(y)

	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

func playWithAppendSlices() {
	var x, y []int

	for i := 0; i < 10; i++ {
		y = appendSlice(x, i)
		fmt.Printf(" %d Capacity=%d \t %v \n", i, cap(y), y)
		x = y
	}

	fmt.Println("\n", x)
	x = appendSlice(x, 10, 20, 30)
	fmt.Println(x)
	x = appendSlice(x, 11, 22, 33, 44, 55, 66)
	fmt.Println(x)
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x = appendSlice(x, numbers...)
	fmt.Println(x)
}

func sliceAppentFromFirst(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func sliceAppendFromLast(strings []string) []string {
	out := strings[:0] // Zero Length Slice Of Original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func playWithNonEmpty() {
	data := []string{"One", "", "Three", "Four", "", "Nice"}
	fmt.Printf("%q \n", data)
	filteredData := sliceAppentFromFirst(data)
	fmt.Printf("%q \n", filteredData)

	dataAgain := []string{"One", "", "Three", "Four", "", "Nice"}
	fmt.Printf("%q \n", dataAgain)
	filteredDataAgain := sliceAppendFromLast(dataAgain)
	fmt.Printf("%q \n", filteredDataAgain)
}

func playWithSlices() {
	fmt.Println("Function: playWithSlice")
	playWithSlice()
	fmt.Println("\n---------------->")

	fmt.Println("Function: playWithSliceComparision")
	playWithSliceComparision()
	fmt.Println("\n---------------->")

	fmt.Println("Function: playWithSliceCapasity")
	playWithSliceCap()
	fmt.Println("\n---------------->")

	fmt.Println("Function: append() behind the sceen")
	playWithAppendInt()
	fmt.Println("\n---------------->")
	// OR
	fmt.Println("Function: append() behind the sceen")
	playWithAppendSlices()
	fmt.Println("\n---------------->")

	fmt.Println("Function: playWithNonEmpty")
	playWithNonEmpty()
	fmt.Println("\n---------------->")
}
