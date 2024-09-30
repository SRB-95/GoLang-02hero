package functions

import (
	"fmt"
	"log"
)

func sumInts(args map[string]int64) int64 {
	var s int64
	for _, v := range args {
		s += v
	}
	return s
}

func sumFloats(args map[string]float64) float64 {
	var s float64
	for _, v := range args {
		s += v
	}
	return s
}

// here K & V are type placeholder
func sumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// By using K comparable, the function can accept maps with keys of any type that supports
// comparison operations (== and !=), not just strings. This makes the function more versatile and reusable.

func playWithGenerics() {
	someInts := map[string]int64{
		"first":  34,
		"second": 12,
		"third":  10,
	}

	someFloats := map[string]float64{
		"first":  34.1,
		"second": 12.2,
		"third":  10.3,
	}

	fmt.Printf("Sum of Ints:  %v\n", sumInts(someInts))
	fmt.Printf("Sum of Ints:  %v\n", sumFloats(someFloats))

	fmt.Printf("Sum of Ints:  %v\n", sumIntsOrFloats[string, int64](someInts))
	fmt.Printf("Sum of Ints:  %v\n", sumIntsOrFloats[string, float64](someFloats))

}

func PrintS(s []string) {
	for _, v := range s {
		fmt.Print(v)
	}
	fmt.Println()
}

func PrintI(s []int) {
	for _, v := range s {
		fmt.Print(v)
	}
	fmt.Println()
}

func Print[T any](args []T) {
	for _, v := range args {
		fmt.Print(v)
	}
	fmt.Println()
}

func playWithGenericPrint() {
	PrintS([]string{"Ding", "Dong", "Ting", "Tong"})
	PrintI([]int{10, 20, 30, 22, 99, 77})

	Print([]string{"Ding", "Dong", "Ting", "Tong"})
	Print([]int{10, 20, 30, 22, 99, 77})
}

func Equal[T comparable](a, b T) bool {
	return a == b
}

func playWithGenericEqual() {
	fmt.Println(Equal("Ding", "Ding"))
	fmt.Println(Equal("Ding", "Dong"))

	fmt.Println(Equal(99, 99))
	fmt.Println(Equal(99, 100))
}

type PrintNum int
type PrintText string

type PrintInterface interface {
	print()
}

func (p PrintNum) print() {
	log.Println("Num:", p)
}

func (p PrintText) print() {
	log.Println("Text:", p)
}

// PrintSlice prints elements of a slice of any type
func PrintSlice[T PrintInterface](s []T) {
	for _, value := range s {
		value.print()
	}
}

func playWithGenericPrintSlice() {
	// Example with PrintSlice
	intSlice := []PrintNum{5, 2, 9}
	stringSlice := []PrintText{"apple", "banana", "orange"}

	fmt.Println("Integer slice:")
	PrintSlice(intSlice)

	fmt.Println("\nString slice:")
	PrintSlice(stringSlice)

}

func playWithGeneric() {
	fmt.Println("function: playWithGenerics")
	playWithGenerics()
	fmt.Println("\n---------------->")

	fmt.Println("function: playWithGenericPrint")
	playWithGenericPrint()
	fmt.Println("\n---------------->")

	fmt.Println("function: playWithGenericEqual")
	playWithGenericEqual()
	fmt.Println("\n---------------->")

	fmt.Println("function: playWithGenericPrintSlice")
	playWithGenericPrintSlice()
	fmt.Println("\n---------------->")

	// fmt.Println("function: ")
	// fmt.Println("\n---------------->")

	// fmt.Println("function: ")
	// fmt.Println("\n---------------->")

	// fmt.Println("function: ")
	// fmt.Println("\n---------------->")
}
