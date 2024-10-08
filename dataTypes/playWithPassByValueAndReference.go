package main

import "fmt"

func modifyArray(arr [3]int) {
	arr[0] = 100 // modifies only the local copy of arr
	fmt.Println("Modify Arr: ", arr)

}

func arryPassByValue() {
	arr := [3]int{1, 2, 3}
	fmt.Println("Original Arr: ", arr)
	modifyArray(arr)
}

func modifyArray1(arr *[3]int) {
	(*arr)[0] = 100 // modifies the original array
	fmt.Println("Modify Arr: ", arr)
}

func arrayByReference() {
	arr := [3]int{10, 20, 30}
	fmt.Println("Original Arr: ", arr)
	modifyArray1(&arr) // pass the address of the array
}

func modifySlice(s *[3]int) {
	(*s)[0] = 100 // modifies the original array
	fmt.Println("Modify Slice: ", s)
}

func sliceByReference() {
	slice := [3]int{10, 20, 30}
	fmt.Println("Original Slice: ", slice)
	modifySlice(&slice) // pass the address of the array
}

func modifyMap(m map[string]int) {
	m["one"] = 100
	fmt.Println("Modify Map: ", m)
}

func mapByReference() {
	map1 := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("Original Map: ", map1)
	modifyMap(map1)
}

type Person struct {
	name string
	age  int
}

func modifyPerson(p Person) {
	p.age = 30
	fmt.Println("Modify Struct: ", p)
}

func structPassByValue() {
	person := Person{name: "Alice", age: 25}
	fmt.Println("Original Struct: ", person.age)
	modifyPerson(person)
}

func modifyPersonAgain(p *Person) {
	p.age = 30                         // modifies the original struct via pointer
	fmt.Println("Modify Struct: ", p)  //  printing a pointer to a struct will display the memory address
	fmt.Println("Modify Struct: ", *p) // Dereference the pointer
}

func structPassByReference() {
	person := Person{name: "John", age: 25}
	fmt.Println("Original Struct: ", person)
	modifyPersonAgain(&person) // pass the address of the struct
}

func main() {
	// array: pass by value
	fmt.Println("Function: Array pass by value")
	arryPassByValue()
	fmt.Println("---------------->\n")

	fmt.Println("Function: Array pass by reference")
	arrayByReference()
	fmt.Println("---------------->\n")

	// silces: pass by reference
	fmt.Println("Function: Slice pass by reference")
	sliceByReference()
	fmt.Println("---------------->\n")

	// Map: pass by reference
	fmt.Println("Function: Map pass by reference")
	mapByReference()
	fmt.Println("---------------->\n")

	// struct: pass by value
	fmt.Println("Function: Struct pass by value")
	structPassByValue()
	fmt.Println("---------------->\n")

	// struct: pass by reference
	fmt.Println("Function: Struct pass by reference")
	structPassByReference()
	fmt.Println("---------------->\n")
}
