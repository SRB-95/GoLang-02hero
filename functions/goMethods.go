package functions

import "fmt"

// Defining a struct type
type Person struct {
	Name string
	Age  int
}

// Method with a value receiver
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s.\n", p.Name)
}

// Method with a pointer receiver
func (p *Person) HaveBirthday() {
	p.Age++
}

func exploreMethods() {
	// Creating an instance of the Person struct
	p := Person{Name: "Alice", Age: 25}

	// Calling a method with a value receiver
	p.Greet()

	// Calling a method with a pointer receiver
	p.HaveBirthday()

	// Checking if age incremented
	fmt.Printf("%s is now %d years old.\n", p.Name, p.Age)
}

func playWithMethods() {
	fmt.Println("Function: Play with Methods")
	exploreMethods()
	fmt.Println("------------------>")

}
