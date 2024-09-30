package functions

import "fmt"

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
}
