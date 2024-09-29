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
	fmt.Println("-----------------------------------: Hell0 Go :------------------------------------>")

	fmt.Println("function: PlayWithDataTypes")
	dataTypes.PlayWithDataTypes()
	fmt.Println("-----------------------------------: Data Types :------------------------------------>")

	fmt.Println("function: playWithAssignments")
	// playWithAssignments()
	fmt.Println("----------------------------------------------------")
}
