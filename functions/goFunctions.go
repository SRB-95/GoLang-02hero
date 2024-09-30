package functions

import (
	"fmt"
	"math"
)

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

//________________________________________________________________

func division(dividend, divisor int) (int, int) {
	if divisor != 0 {
		quotient := dividend / divisor
		remainder := dividend % divisor
		return quotient, remainder
	} else {
		fmt.Println("Error: Division By Zero!")
	}
	return 0, 0
}

func divisionAgain(dividend, divisor int) (quotient int, remainder int) {
	if divisor != 0 {
		quotient := dividend / divisor
		remainder := dividend % divisor
		return quotient, remainder
	} else {
		fmt.Println("Error: Division By Zero!")
	}
	return 0, 0
}

func playWithDivision() {
	q, r := division(10, 3)
	fmt.Printf("\nQuotient: %d Remainder: %d", q, r)

	qq, rr := divisionAgain(10, 3)
	fmt.Printf("\nQuotient: %d Remainder: %d", qq, rr)
}

func playWithClosures() {
	var x, y = 40, 20

	// Anonymous/Closure/Lambda Function
	add := func(x, y int) int { return x + y }
	sub := func(x, y int) int { return x - y }

	result := add(x, y)
	fmt.Println("Addition: ", result)

	result = sub(x, y)
	fmt.Println("Subtraction: ", result)
}

func sum(x, y int) int { return x + y }
func sub(x, y int) int { return x - y }
func mul(x, y int) int { return x * y }

// Polymorphic function
//    Takes multiple forms
//    Using mechanise: By passing function/Closure to function

func calculator(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func playWithCalculator() {
	var a, b = 40, 20
	var result = 0

	result = calculator(a, b, sum)
	fmt.Println("Addition: ", result)

	result = calculator(a, b, sub)
	fmt.Println("Subtraction: ", result)

	result = calculator(a, b, mul)
	fmt.Println("Multiplication: ", result)

	sumClosure := func(x, y int) int { return x + y }
	subClosure := func(x, y int) int { return x - y }
	mulClosure := func(x, y int) int { return x * y }

	result = calculator(a, b, sumClosure)
	fmt.Println("Addition: ", result)

	result = calculator(a, b, subClosure)
	fmt.Println("Subtraction: ", result)

	result = calculator(a, b, mulClosure)
	fmt.Println("Multiplication: ", result)
}
func Sphere() func(radius float64) float64 {
	volumeClosure := func(radius float64) float64 {
		volume := (4.0 / 3.0) * math.Pi * math.Pow(radius, 3)
		return volume
	}
	return volumeClosure
}

func playWithHigherOrderFunctionsAgain() {
	// Define a function that returns a closure to compute the volume
	something := Sphere()

	// Invoking the closure
	fmt.Println("Volume Of Sphere with radius 5: ", something(5))
	fmt.Println("Volume Of Sphere with radius 10: ", something(10))

	// Assign Sphere function itself (without calling it)
	somethingAgain := Sphere
	// Print its type, not the function object
	fmt.Printf("\nType Of somethingAgain: %T\n", somethingAgain)

	// Print the type of Sphere function
	fmt.Printf("Type Of Sphere: %T\n", Sphere)

	// Print the type of Sphere function reference
	fmt.Printf("\nType Of Sphere (as reference): %T\n", Sphere)
}

func PlayWithFunctions() {
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

	fmt.Println("Function: Generic functions")
	playWithGeneric()
	fmt.Println("\n---------------->")

	fmt.Println("Function: playWithDivision")
	playWithDivision()
	fmt.Println("------------------>")

	fmt.Println("Function: playWithClosures")
	playWithClosures()
	fmt.Println("------------------>")

	fmt.Println("Function: playWithCalculator")
	playWithCalculator()
	fmt.Println("------------------>")

	fmt.Println("Function: playWithHighorderFunctionsAgain")
	playWithHigherOrderFunctionsAgain()
	fmt.Println("------------------->")

	// Methods
	fmt.Println("Function: playWithMethods")
	playWithMethods()
	fmt.Println("------------------>")
}
