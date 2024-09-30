package dataTypes

import "fmt"

func modifyValue(ptr *int) {
	*ptr = 50 // Modify the value via the pointer
}

func failedUpdate(g *int) {
	x := 10
	g = &x
}

func playWithPointers() {
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
