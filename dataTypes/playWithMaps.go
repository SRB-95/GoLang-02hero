package dataTypes

import (
	"bufio"
	"fmt"
	"os"
)

// Maps Is Collection Of Key and Value Pairs
// Map Syntax : map[ KeyType ]ValueType

func exploreMaps() {
	// way1: Create map using map()
	m := make(map[string]int)

	m["k1"] = 7  // Storing 7 Value For Key "k1"
	m["k2"] = 13 // Storing 13 Value For Key "k2"

	fmt.Println("map:", m)

	v1 := m["k1"] // Getting Value For Key "k1"
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2") // Deleting Key "k2" : Will Delete Key, Value Pair
	fmt.Println("map:", m)

	_, prs := m["k2"] // Getting Value For Key "k2"
	fmt.Println("prs:", prs)

	// way2: To Create Map With Initialisation List
	// Creating and initializing a map Using shorthand declaration and using map literals
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

func removeDuplicates() {
	seen := make(map[string]bool) // a set of strings
	fmt.Println("Length of seen: ", len(seen))
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "removeDuplicates: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(seen)
}

// create and initialize maps
func isMapNil() {
	// way3: Creating and initializing empty map Using var keyword
	var map_1 map[int]int

	// Checking if the map is nil or not
	if map_1 == nil {
		fmt.Println("isMapNil: ", "True")
	} else {
		fmt.Println("isMapNil: ", "False")
	}

	map_2 := map[int]string{
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}

	fmt.Println("Map-2: ", map_2)
}

func playWithMaps2() {
	// Creating a map
	// Using make() function
	var My_map = make(map[float64]string)
	fmt.Println("My_map: ", My_map)

	// As we already know that make() function
	// always returns a map which is initialized
	// So, we can add values in it
	My_map[1.3] = "Rohit"
	My_map[1.5] = "Sumit"
	fmt.Println("My_map:->", My_map)
}

//________________________________________________________________

// To iterate the map using for range loop
func playWithMaps3() {
	// Creating and initializing a map
	m_a_p := map[int]string{
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}

	// Iterating map using for range loop
	for id, pet := range m_a_p {
		fmt.Println("m_a_p: ", id, pet)
	}
}

//________________________________________________________________

// a key-value pair in the map using make() function
func playWithMaps4() {
	// Creating and initializing a map
	m_a_p := map[int]string{
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}

	fmt.Println("Original map: ", m_a_p)

	// Adding new key-value pairs in the map
	m_a_p[95] = "Parrot"
	m_a_p[96] = "Crow"
	fmt.Println("Map after adding new key-value pair:\n", m_a_p)

	// Updating values of the map
	m_a_p[91] = "PIG"
	m_a_p[93] = "DONKEY"
	fmt.Println("\nMap after updating values of the map:\n", m_a_p)
}

//________________________________________________________________

// retrieve the value of the key
func playWithMaps5() {
	// Creating and initializing a map
	m_a_p := map[int]string{
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}
	fmt.Println("Original map: ", m_a_p)

	// Retrieving values with the help of keys
	value_1 := m_a_p[90]
	value_2 := m_a_p[93]
	fmt.Println("Value of key[90]: ", value_1)
	fmt.Println("Value of key[93]: ", value_2)
}

//________________________________________________________________

// check the key is available or not
func playWithMaps6() {
	// Creating and initializing a map
	m_a_p := map[int]string{
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}

	fmt.Println("Original map: ", m_a_p)

	// Checking the key is available
	// or not in the m_a_p map
	pet_name, ok := m_a_p[90]
	fmt.Println("\nKey present or not:", ok)
	fmt.Println("Value:", pet_name)

	// Using blank identifier
	_, ok1 := m_a_p[92]
	fmt.Println("\nKey present or not:", ok1)
}

//________________________________________________________________

// Go program to illustrate how to delete a key
func playWithMaps7() {
	// Creating and initializing a map
	m_a_p := map[int]string{
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}

	fmt.Println("Original map: ", m_a_p)

	// Deleting keys
	// Using delete function
	delete(m_a_p, 90)
	delete(m_a_p, 93)

	fmt.Println("Map after deletion: ", m_a_p)
}

//________________________________________________________________

// modification concept in map
func playWithMaps8() {
	// Creating and initializing a map
	m_a_p := map[int]string{
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}
	fmt.Println("Original map: ", m_a_p)

	// Assigned the map into a new variable
	new_map := m_a_p

	// Perform modification in new_map
	new_map[96] = "Parrot"
	new_map[98] = "Pig"

	// Display after modification
	fmt.Println("New map: ", new_map)
	fmt.Println("\nModification done in old map:\n", m_a_p)
}

func playWithMaps() {
	fmt.Println("Function: exploreMaps")
	exploreMaps()
	fmt.Println("\n---------------->")

	fmt.Println("Function: removeDuplicates")
	removeDuplicates()
	fmt.Println("\n---------------->")

	fmt.Println("Function: isMapNil")
	isMapNil()
	fmt.Println("\n---------------->")

	playWithMaps2()
	playWithMaps3()
	playWithMaps4()
	playWithMaps5()
	playWithMaps6()
	playWithMaps7()
	playWithMaps8()
}
