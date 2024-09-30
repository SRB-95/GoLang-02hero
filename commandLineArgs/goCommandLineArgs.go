package commandLineArgs

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func playWithCommandLineArguements() {
	var someString, separator string
	separator = "\n"
	var args = os.Args
	for i := 0; i < len(args); i++ {
		someString = someString + separator + args[i]
	}
}

func playWithReadingInputs() {
	fmt.Println("Please enter multiple integer devided by space...!, After that press Enter")
	input := bufio.NewScanner(os.Stdin) // reading inputs from terminal

outer:
	for input.Scan() {
		var ints []int

		for _, s := range strings.Fields(input.Text()) { // strings.Fields(input.Text() => [ram ram ram]
			x, err := strconv.ParseInt(s, 10, 64)

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}

		fmt.Printf("Inputs are: %v\n", ints)
		slices.Reverse(ints)
		fmt.Printf("After Reversing: %v\n", ints)
		fmt.Println("Please press ctrl+D to exit")
	}
}

func countUnicodeCharacters() {
	counts := make(map[rune]int)    // counts of Unicode characters [Rune: U+0041]
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	fmt.Println("Please enter multiple words devided by line...!")
	fmt.Println("Press Enter and then ctrl+D to exit...!")
	in := bufio.NewReader(os.Stdin) // bufio are useful when reading inputs line by line

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {         // Ensure the ends of file
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

// Binary Tree Node
type tree struct {
	value int
	left  *tree
	right *tree
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func playWithTree() {
	data := []int{10, 20, 30, 9, 80, 15, 14, 16}

	var root *tree
	for _, value := range data {
		root = add(root, value)
	}

	fmt.Println(data)

	var values []int
	fmt.Println(appendValues(values, root))
}

type Point struct {
	X, Y int
}

func ScalePoint(point Point, factor int) Point {
	var x, y int
	x = point.X * factor
	y = point.Y * factor
	return Point{x, y}
}

func AddPoint(point1 Point, point2 Point) Point {
	var x, y int
	x = point1.X + point2.X
	y = point1.Y + point2.Y
	return Point{x, y}
}

func playWithPointType() {
	// Creating Point Type Objects
	// 	Using Declarative Syntax With Initialisation List
	point1 := Point{10, 20}
	point2 := Point{100, 200}
	point11 := Point{10, 20}

	fmt.Println("point1 : ", point1)
	fmt.Println("point2 : ", point2)

	point3 := ScalePoint(point1, 5)
	point4 := ScalePoint(point2, 5)
	fmt.Println("point3 = point1 * 5 : ", point3)
	fmt.Println("point4 = point2 * 5 : ", point4)

	point5 := AddPoint(point1, point2)
	fmt.Println("point1 + point2 : ", point5)

	// HOME WORK
	fmt.Println("Points Equals :", point1.X == point11.X && point1.Y == point11.Y)
	fmt.Println("Points Equals :", point1 == point11)
	fmt.Println("Points Equals :", point1 == point2)
}

func PlayWithCommandLineArgs() {
	fmt.Println("Function: Command line arguements")
	playWithCommandLineArguements()
	fmt.Println("\n---------------->")

	fmt.Println("Function: playWithReadingInputs")
	playWithReadingInputs()
	fmt.Println("\n---------------->")

	fmt.Println("Function: countUnicodeCharacters")
	countUnicodeCharacters()
	fmt.Println("\n---------------->")

	fmt.Println("Function: playWithTree")
	playWithTree()
	fmt.Println("\n---------------->")

	fmt.Println("Function: playWithPointType")
	playWithPointType()
	fmt.Println("\n---------------->")
}
