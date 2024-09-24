package dataTypes

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Constants
const boilingf = 212.0

var ff = 99
var cc = 88

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func playWithVarAndConst() {
	var f = boilingf
	var c = (f - 32) * 5 / 9
	fmt.Printf("\nBoiling Point: %g Farenheit", f)
	fmt.Printf("\nBoiling Point: %g Centiigrade", c)
	var boilingC = fToC(boilingf)
	fmt.Printf("\nBoiling Point:-> %g Centiigrade", boilingC)

	const first, second float64 = 99.99, 88.88             // Multiple variable declaration
	fmt.Printf("\nFirst: %g, Second: %g", first, second)   // Format specifier %g removes extra 0 i.e 99.9000
	fmt.Printf("\nAccess Global variables: %d %d", ff, cc) // Format specifier %d used to format integer
}

type Celsius float64
type Farenheit float64

const (
	Boilingc     Celsius   = 100.0
	FreezingC    Celsius   = 0.0
	AbsoluteZerC Celsius   = -273.15
	FreezingF    Farenheit = 0.05
)

func playWithTypes() {
	// Integer
	var a int
	var a8 int8   // -128 to 127
	var a16 int16 // -32,768 to 32,767
	var a32 int32 // -2,147,483,648 to 2,147,483,647
	var a64 int64 //  (-9,223,372,036,854,775,808 to 9,223,372,036,854,775,807)
	fmt.Println("Default Values-int: ", a, a8, a16, a32, a64)

	var ua uint
	var ua8 uint8   // 0 to 255
	var ua16 uint16 // 0 to 65,535
	var ua32 uint32 // (0 to 4,294,967,295)
	var ua64 uint64 // (0 to 18,446,744,073,709,551,615)
	fmt.Println("Default Values-uint: ", ua, ua8, ua16, ua32, ua64)

	// Float
	var f1 float32
	var f2 float64
	fmt.Println("Default values-float: ", f1, f2)

	// String
	var some string
	fmt.Println("Default value-string: ", some)
	some = "Go lang"
	fmt.Println("Re-assigned some: ", some)
	fmt.Printf("Type of some: %T", some) // %T used for type

	// Boolean
	var b bool
	fmt.Println("Default values-b: ", b)

	// Complex
	var c1 complex64
	var c2 complex128
	fmt.Println("Default values-c1, c2: ", c1, c2)

	var ff, gg = 11.11, 22.22
	var rr = ff + gg
	fmt.Println("\nResult: ", rr)

	var ffAgain Celsius = 11.11
	var ggAgain Celsius = 22.22 // Farenheit
	fmt.Println("Values Again: ", ffAgain, ggAgain)
	var rrAgain = ffAgain + ggAgain // if ggAgain is type of Farenheit, then it will through error: mismatched types Celsius and Farenheit
	fmt.Println("ResultAgain: ", rrAgain)

	var someInt int64 = 99
	var someFloat float64 = 11.11
	var something = someInt + int64(someFloat) // Explictly type changed
	fmt.Println(something)
}

func playWithIntegers() {
	a := 10
	b := 3
	sum := a + b // 13
	fmt.Printf("Sum: %d", sum)
	difference := a - b // 7
	fmt.Printf("\nDifference: %d", difference)
	product := a * b // 30
	fmt.Printf("\nProduct: %d", product)
	division := a / b // 3
	fmt.Printf("\nDivision: %d", division)
	remainder := a % b // 1
	fmt.Printf("\nRemainder: %d", remainder)

	fmt.Println("\n---------------->")

	and := a & b
	fmt.Printf("\nand value: %d", and)
	or := a | b
	fmt.Printf("\nor value: %d", or)
	xor := a ^ b
	fmt.Printf("\nxor value: %d", xor)
	not := ^a
	fmt.Printf("\nnot value: %d", not)
	shiftLeft := a << 1
	fmt.Printf("\nshiftLeft value: %d", shiftLeft)
	shiftRight := a >> 1
	fmt.Printf("\nshiftRight value: %d", shiftRight)

	fmt.Println("\n---------------->")

	// integer to string conversion
	num := 42
	str := strconv.Itoa(num) // "42"
	fmt.Println("int of str: ", str)

	// string to integer conversion
	num, err := strconv.Atoi(str) // 42, error
	fmt.Println("str to int: ", num, err)
}

func playWithStrings() {
	word := "golang"
	fmt.Println("Length of words: ", len(word))
	char := word[2]
	fmt.Printf("Character: %c", char)
	word2 := " is beautiful"
	result := word + word2
	fmt.Println("\nConcatination: ", result)
	isContains := strings.Contains(result, "beautiful")
	fmt.Println("String Containg: ", isContains)
	containsAny := strings.ContainsAny(word, "aeiou")
	fmt.Println("Strings Contains any: ", containsAny)
	prefix := strings.HasPrefix(word, "go")
	fmt.Println("String has prefix: ", prefix)
	suffix := strings.HasSuffix(word, "ng")
	fmt.Println("String has suffix: ", suffix)
	index := strings.Index(word, "o")
	fmt.Println("Index of a char in word: ", index)
	lastIndex := strings.LastIndex(word, "l")
	fmt.Println("Last index of word: ", lastIndex)
	upperCase := strings.ToUpper(word)
	fmt.Println("Uppercase: ", upperCase)
	lowerCase := strings.ToLower(upperCase)
	fmt.Println("lowercase: ", lowerCase)
	split := strings.Split(result, " ")
	fmt.Println("Split: ", split)
	join := strings.Join(split, "-")
	fmt.Println("Join: ", join)
	trim := strings.Trim("...hello...", ".") // "hello"
	fmt.Println("stringTrim: ", trim)
	trimSpace := strings.TrimSpace("   hello   ") // "hello"
	fmt.Println("trim space: ", trimSpace)
	replace := strings.Replace("hello world", "l", "x", 2)
	fmt.Println("string replace: ", replace)
}

func playWithBoolean() {
	launch := false
	launchText := fmt.Sprintf("%v", launch)
	fmt.Println("Ready for launch:", launchText)
}

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
	// Array declare & initialization
	var arr1 [3]int // [0 0 0], An array is declared with a size (e.g., [3]int), and its size is part of its type.
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	fmt.Println("Arr1: ", arr1)
	fmt.Println("\n---------------->")
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
	someArr := [...]int{0: 100, 2: 200, 300}
	fmt.Println("Controll on index of an array: ", someArr)
	fmt.Println("Length of an array: ", len(someArr))
	fmt.Printf("Type of an array: %T", someArr)
	fmt.Println("\n---------------->")
	// in Go, both arrays and their elements are passed by value.
	// This means that when you pass an array to a function, a copy of the entire array is made, and changes to this copy do not affect the original array.
	// Similarly, when you access or modify an array element, you are working with a copy of the value unless you use pointers.
	arr4 := [3]int{9, 99, 999} // Original array
	modifyArray(arr4)          // Passing array to the function
	fmt.Println("Outside modifyArray:", arr4)
	fmt.Println("\n---------------->")
	// If you want to modify the original array or elements, you need to pass a pointer to the array or its elements:
	arr5 := [3]int{1, 2, 3}

	// Passing the array by reference
	modifyArrayByReference(&arr5)
	fmt.Println("Outside modifyArrayByReference:", arr5) // Now the array is modified

	// Passing the array element by reference
	modifyElementByReference(&arr5[1])
	fmt.Println("Outside modifyElementByReference:", arr5[1]) // The specific element is modified
}

func playWithSlices() {
	slice := make([]int, 3, 5) // By using make() function we can creates a slice of length 3 and capacity
	fmt.Println(slice)         // Output: [0 0 0] (initialized with zero values)
	fmt.Println("Length of slice: ", len(slice))
	fmt.Println("Capacity of slice: ", cap(slice))
	slice = append(slice, 4, 5) //  Adds new elements to the slice. If the slice's capacity is not enough.
	// When length & capasity becomes same, then automaically capacity besomes doubled
	fmt.Println(slice)
	copyArr := []int{10, 20, 30, 40, 50}
	copy(copyArr, slice) // deep copy
	fmt.Println("Copy Arr: ", copyArr)
	fmt.Println("\n---------------->")
	slice1 := [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	some1 := slice1[0:4]
	fmt.Println(some1)
	some2 := slice1[4:9]
	fmt.Println(some2)
	some3 := slice1[:5]
	fmt.Println(some3)
	some4 := slice1[6:]
	fmt.Println(some4)
	some5 := slice1[:]
	fmt.Println(some5)
}

func modifyValue(ptr *int) {
	*ptr = 50 // Modify the value via the pointer
}

func playWithPointer() {
	var name string = "Happy"
	fmt.Println("Memory Address: ", &name) // Memory Address:  0xc0000140a0

	// create the pointer variable
	// var ptr *string // ptr declared but not initialize so the val is nil
	// ptr = &name
	// var ptr = new(int) // Also we can create a ptr using new()
	var ptr *string = &name // assign the memory address of name to the pointer
	// var ptr = &name //Also we can create a ptr w/o using *
	fmt.Println("Value of pointer is", ptr)
	fmt.Println("Address of the variable", &name)

	fmt.Println(*ptr) // * is used to get the value pointed by ptr
	*ptr = "Soumya"   // Dereferencing and Modifying Values
	fmt.Println(*ptr)
	fmt.Println("\n---------------->")
	// In Go, when you pass variables to functions, they are passed by value by default,
	// meaning the function receives a copy of the variable. However, if you pass a pointer to a variable,
	// the function can modify the original value (pass by reference).
	num := 10
	fmt.Println("Before:", num) // Output: 10
	modifyValue(&num)           // Pass the address of num to the function
	fmt.Println("After:", num)  // Output: 50
}

func changeIt(ptr *int) {
	*ptr = 0
}

func calcAge(name string) func(int) {
	//  Anonymus/literal function
	calcAge := func(yob int) { // calcAge function calculates and prints the age
		fmt.Printf("I am %s, %d years old.\n", name, 2023-yob)
	}
	return calcAge
}

func countDown(num int) {
	fmt.Println(num)
	updatedNum := num - 1
	if updatedNum > 0 {
		countDown(updatedNum) // Recursive call
	}
}

func variadic(args ...int) int { // (...)ellipsis
	total := 0
	for _, val := range args {
		total += val
	}
	fmt.Println("Total: ", total)
	return total
}

func riskyDivision(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()

	if b == 0 {
		panic("division by zero")
	} else {
		fmt.Println("Result:", a/b)
	}
}

func playWithFunctions() {
	fmt.Println("Function pass by reference")
	x := 5
	fmt.Println("First", x)
	changeIt(&x)
	fmt.Println("Then", x)
	fmt.Println("\n---------------->")
	fmt.Println("closure function example")
	greet := calcAge("John")
	greet(1994)
	fmt.Println("\n---------------->")
	fmt.Println("recursive function example")
	countDown(5)
	fmt.Println("\n---------------->")
	fmt.Println("Variadic Function")
	variadic(1, 2, 3)
	fmt.Println("\n---------------->")
	fmt.Println("Function: playWithDeferPanicRecover")
	fmt.Println("Start")
	riskyDivision(10, 0)
	fmt.Println("End")
}

func playWithCommandLineArguements() {
	var someString, separator string
	separator = "\n"
	var args = os.Args
	for i := 0; i < len(args); i++ {
		someString = someString + separator + args[i]
	}
}

func PlayWithDataTypes() {
	fmt.Println("function: playWithVarAndConst")
	playWithVarAndConst()
	fmt.Println("\n----------------------------------------------------")
	fmt.Println("function: playWithTypes")
	playWithTypes()
	fmt.Println("----------------------------------------------------")
	fmt.Println("function: playWithIntegers")
	playWithIntegers()
	fmt.Println("----------------------------------------------------")
	fmt.Println("function: playWithStrings")
	playWithStrings()
	fmt.Println("----------------------------------------------------")
	fmt.Println("function: playWithBoolean")
	playWithBoolean()
	fmt.Println("----------------------------------------------------")
	fmt.Println("function: playWithArrays")
	playWithArrays()
	fmt.Println("----------------------------------------------------")
	fmt.Println("function: playWithSlices")
	playWithSlices()
	fmt.Println("----------------------------------------------------")
	fmt.Println("function: playWithPointer")
	playWithPointer()
	fmt.Println("----------------------------------------------------")
	fmt.Println("function: playWithFunctions")
	playWithFunctions()
	fmt.Println("----------------------------------------------------")
	fmt.Println("function: playWithCommandLineArguements")
	playWithCommandLineArguements()
	fmt.Println("----------------------------------------------------")
}
