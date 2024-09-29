package dataTypes

import "fmt"

func modifyArray(arr [3]int) {
	arr[0] = 10                             // Modify the first element
	fmt.Println("Inside modifyArray:", arr) // Prints the modified copy
}

// Function to modify the array using a pointer
func modifyArrayByReference(arr *[3]int) {
	arr[0] = 100 // Modify the first element
	fmt.Println("Inside modifyArrayByReference:", arr)
}

// Function to modify an array element using a pointer
func modifyElementByReference(elem *int) {
	*elem = 200 // Modify the element via pointer
	fmt.Println("Inside modifyElementByReference:", *elem)
}

func playWithArrays() {
	// Way1: Array declaration & initialization
	var arr1 [3]int // [0 0 0], An array is declared with a size (e.g., [3]int), and its size is part of its type.
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	fmt.Println("Arr1: ", arr1)
	fmt.Println("\n---------------->")

	// Way2-Array literals: Array declaration & initialization
	arr2 := [...]int{1, 2, 3, 4}
	fmt.Println("Arr2: ", arr2)
	fmt.Println("\n---------------->")

	arr3 := [5]int{1, 2, 3, 4, 5}
	for _, ele := range arr3 { // if index is not in used can be replace with _
		fmt.Println(ele)
	}
	fmt.Println("Arr3 first element: ", arr3[0])
	fmt.Println("Arr3 element range: ", arr3[:3]) // [1 2 3]
	fmt.Println("Arr3 element range: ", arr3[3:]) // [4 5]
	copyArr := arr3                               // Copying aray
	copyArr[0] = 10                               // modifying arr3 does not affect arr1
	fmt.Println(copyArr)                          // [10 2 3 4 5]
	fmt.Println(len(arr3), arr3)                  // [1 2 3 4 5]
	// in Go array is passed by value. This means that when you pass an array to a function, Go creates a copy of the array, and any changes made to the array inside the function do not affect the original array.
	fmt.Println("\n---------------->")

	aa := [2]int{10, 20}
	bb := [...]int{10, 20}
	fmt.Println("aa is Equal to bb: ", aa == bb) // true
	// In Go lang first will comapre the type, if matches then will compare the value or any operation will happen
	fmt.Println("\n---------------->")

	someArr := [...]int{0: 100, 2: 200, 300} // We can control over the index of an array
	fmt.Println("Controll on index of an array: ", someArr)
	fmt.Println("Length of an array: ", len(someArr))
	fmt.Printf("Type of an array: %T", someArr)
	fmt.Println("\n---------------->")

	// in Go, both arrays and their elements are passed by value.
	// This means that when you pass an array to a function, a copy of the entire array is passed, and changes to this copy do not affect the original array.
	// Similarly, when you access or modify an array element, you are working with a copy of the value unless you use pointers.
	arr4 := [3]int{9, 99, 999} // Original array
	modifyArray(arr4)          // Passing array to the function
	fmt.Println("Outside modifyArray:", arr4)
	fmt.Println("\n---------------->")

	// If you want to modify the original array or elements, you need to pass a pointer to the array or its elements:
	arr5 := [3]int{1, 2, 3}

	modifyArrayByReference(&arr5)                        // Passing the array by reference
	fmt.Println("Outside modifyArrayByReference:", arr5) // Now the array is modified

	// Passing the array element by reference
	modifyElementByReference(&arr5[1])
	fmt.Println("Outside modifyElementByReference:", arr5[1]) // The specific element is modified

	// Multi Dimentional Array:
}
