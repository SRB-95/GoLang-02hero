package main

import (
	"fmt"
	"golang-02hero/dataTypes"
)

func helloGo() {
	fmt.Println("Hello Go lang")
}

func main() {
	helloGo()
	fmt.Println("----------------------------------------------------")

	fmt.Println("function: PlayWithDataTypes")
	dataTypes.PlayWithDataTypes()
	fmt.Println("----------------------------------------------------")

	fmt.Println("function: playWithAssignments")
	playWithAssignments()
	fmt.Println("----------------------------------------------------")
}
