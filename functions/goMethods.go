package functions

import (
	"fmt"
	"math"
)

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

type Point3D struct {
	X, Y float64
}

type Path []Point3D

func (this Point3D) Distance(q Point3D) float64 {
	return math.Hypot(q.X-this.X, q.Y-this.Y)
}

func (path Path) Distance() float64 {
	totalDistance := 0.0
	for i := range path {
		if i > 0 {
			totalDistance += path[i-1].Distance(path[i])
		}
	}
	return totalDistance
}

func playWithMethodsAgain() {
	var point1 = Point3D{10.0, 20.0}
	var point2 = Point3D{30.0, 40.0}

	var path Path = Path{point1, point2, {50.0, 60.0}, {70.0, 80.0}}
	distance := path.Distance()
	fmt.Printf("\nDistance between %v and %v : %v", point1, point2, distance)
}

func playWithMethods() {
	fmt.Println("Function: Play with Methods")
	exploreMethods()
	fmt.Println("------------------>")

	fmt.Println("Function: Play with Methods Again")
	playWithMethodsAgain()
	fmt.Println("------------------>")

}
