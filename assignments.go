package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomInt() {
	// Create a new random generator with a source
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random number between 1 and 10
	randomInt := rng.Intn(10) + 1 // Intn(10) gives 0-9, so we add 1 to shift to 1-10
	fmt.Println(randomInt)
}

func leapYear() {
	const year = 2024
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Printf("%d is a leap year.", year)
	} else {
		fmt.Printf("%d is not a leap year.", year)
	}
}

func switchCaase() {
	day := 3

	switch day {
	case 1:
		fmt.Println("\nIt's Monday")
	case 2:
		fmt.Println("\nIt's Tuesday")
	case 3:
		fmt.Println("\nIt's Wednesday")
	case 4:
		fmt.Println("\nIt's Thursday")
	case 5:
		fmt.Println("\nIt's Friday")
	case 6:
		fmt.Println("\nIt's Saturday")
	case 7:
		fmt.Println("\nIt's Sunday")
	default:
		fmt.Println("\nIt's Invalid day")
	}
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr // Base case: arrays with 0 or 1 element are already sorted
	}

	lastElement := arr[0] // Choose the first element as the lastElement
	left := []int{}
	right := []int{}

	// Partition the array into left and right slices
	for _, value := range arr[1:] {
		if value < lastElement {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}

	// Recursively sort the left and right slices, then concatenate
	return append(append(quickSort(left), lastElement), quickSort(right)...)
}

func factorial(num uint) uint {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

// Day6: 25/9/24

func playWithAssignments() {
	getRandomInt()
	fmt.Println("----------------------------------------------------")
	leapYear()
	fmt.Println("\n----------------------------------------------------")
	switchCaase()
	fmt.Println("----------------------------------------------------")
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	sortedArr := quickSort(arr)
	fmt.Println("Sorted array:", sortedArr)
	fmt.Println("----------------------------------------------------")
	fmt.Println("function:  Factorial Number")
	fmt.Println(factorial(5))
	fmt.Println("----------------------------------------------------")
}
