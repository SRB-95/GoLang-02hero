package dataTypes

import (
	"bufio"
	"fmt"
	"os"
)

func exploreMaps() {
	// way1: Create map using map(), which automatically initialise a map
	m := make(map[string]int)
	m["k1"] = 7  // Storing 7 Value For Key "k1"
	m["k2"] = 13 // Storing 13 Value For Key "k2"
	fmt.Println("map:", m)

	v1 := m["k1"] // Getting Value For Key "k1"
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2") // Deleting Key "k2" : Will Delete Key, Value Pair
	fmt.Println("map:", m)

	// Checking the key is available or not in the map
	val, prs := m["k2"]
	fmt.Println("\nKey present or not:", prs)
	fmt.Println("Value: ", val)

	// way2: To Create Map With Initialisation List
	// Creating and initializing a map Using shorthand declaration and using map literals
	m_a_p := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", m_a_p)
	// To iterate the map using for range loop
	for k, v := range m_a_p {
		fmt.Println("m_a_p: ", k, v)
	}
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
	// In Go, when you declare a map using var map_1 map[string]int, it creates a nil map.
	// You cannot add key-value pairs to a nil map without first initializing it using make or a map literal.
	var map_1 map[string]int
	fmt.Println("map_1: ", map_1)

	// Checking if the map is nil or not
	if map_1 == nil {
		fmt.Println("isMapNil: ", "True")
		// Initialize the map using make
		map_1 = make(map[string]int)
		map_1["k1"] = 10
		map_1["k2"] = 20
		fmt.Println("map_1: ", map_1)
		map_1["k1"] = 100
		fmt.Println("After updating map_1: ", map_1["k2"])
		fmt.Println("After updating map_1: ", map_1)
	} else {
		fmt.Println("isMapNil: ", "False")
	}
}

// modification concept in map
func reassigMap() {
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
	new_map[95] = "Parrot"
	new_map[96] = "Pig"

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

	fmt.Println("Function: playWithReassignMap")
	reassigMap()
	fmt.Println("\n---------------->")
}
