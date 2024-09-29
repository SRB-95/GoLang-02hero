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
	fmt.Printf("\nBoiling Point: %g Centigrade", c)
	var boilingC = fToC(boilingf)
	fmt.Printf("\nBoiling Point:-> %g Centigrade", boilingC)

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

func modifyValue(ptr *int) {
	*ptr = 50 // Modify the value via the pointer
}

func failedUpdate(g *int) {
	x := 10
	g = &x
}

func playWithPointer() {
	var name string = "Happy"
	fmt.Println("Memory Address of name: ", &name) // Memory Address:  0xc0000140a0

	// create the pointer variable
	// var ptr *string // ptr declared but not initialize so the val is nil
	// ptr = &name
	var ptr *string = &name // declare & assign the memory address of name to the ptr
	// var ptr = new(int) // Also we can create a ptr using new()
	// fmt.Println("isNil: ", ptr == nil)           // false
	// fmt.Println("Pointer default value: ", *ptr) // 0
	// var ptr = &name //Also we can create a ptr w/o using *
	fmt.Println("Value of pointer is: ", *ptr)

	fmt.Println("Before ptr value: ", *ptr) // * is used to get the value pointed by ptr
	*ptr = "Soumya"                         // Dereferencing and Modifying Values
	fmt.Println("After ptr value: ", *ptr)

	fmt.Println("\n---------------->")

	fmt.Println("You can only reassign the value if there was a value already assigned to the pointer.")
	var f *int // f is nil
	failedUpdate(f)
	fmt.Println(f)

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

func firstFunction() {
	fmt.Println("First function called")
}

func secondFunction() {
	fmt.Println("Second function called")
}

func thirdFunction() {
	fmt.Println("Third function called")
}

func playWithDeferPanicRecover() {
	// defer kew word delays Execution of statemant till End of Enclosing function
	defer firstFunction()
	secondFunction()
	defer thirdFunction()
}

func someFunction() (some int) {
	defer func() { some++ }()
	return 1
}

func platWithTrickyDefer() {
	result := someFunction()
	fmt.Println("Result: ", result)
}

func applyOperation(a, b int, operation func(a, b int) int) int {
	return operation(a, b)
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func playWithHigherOrderFunction() {
	// Passing a Function as an Argument
	sum := applyOperation(3, 5, add)
	fmt.Println("Sum: ", sum)

	mul := applyOperation(3, 5, multiply)
	fmt.Println("Multiplication: ", mul)

	// Returning a Function
	double := createMultiplier(2)
	fmt.Println("Double func: ", double(5))
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

	fmt.Println("Function: playWithDeferPanicRecover") // Last In First Out
	playWithDeferPanicRecover()
	platWithTrickyDefer()
	fmt.Println("\n---------------->")

	fmt.Println("Function: playWithHigherOrderFunction")
	playWithHigherOrderFunction()
	fmt.Println("\n---------------->")

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
	fmt.Println("\n----------------------------------------:")

	fmt.Println("function: playWithTypes")
	playWithTypes()
	fmt.Println("------------------------------------------:")

	fmt.Println("function: playWithIntegers")
	playWithIntegers()
	fmt.Println("------------------------------------------:")

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

	fmt.Println("function: playWithMaps")
	playWithMaps()
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
