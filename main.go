package main

import (
	"fmt"
	"golang-02hero/dataTypes"
	"golang-02hero/functions"
	"golang-02hero/goRoutines"
	"golang-02hero/routes"
)

func helloGo() {
	fmt.Println("Hello Go lang")
}

func main() {
	helloGo()
	fmt.Println("-----------------------------------: Hell0 Go :------------------------------------>")

	fmt.Println("function: PlayWithDataTypes")
	dataTypes.PlayWithDataTypes()
	fmt.Println("---------------------------------: Data Types :------------------------------------>")

	fmt.Println("function: PlayWithFunctions")
	functions.PlayWithFunctions()
	fmt.Println("-----------------------------------: Functions :----------------------------------->")

	fmt.Println("function: Play With Command Line Arguements")
	// commandLineArgs.PlayWithCommandLineArgs()
	fmt.Println("-------------------------------: Command Line Args :------------------------------->")

	fmt.Println("function: Play With Go Routines")
	goRoutines.PlayWithGoRoutines()
	fmt.Println("----------------------------------: Go Routines :---------------------------------->")

	fmt.Println("function: Play With Routing")
	routes.PlayWithRouting()
	fmt.Println("------------------------------------: Routing :------------------------------------>")

	fmt.Println("function: playWithAssignments")
	playWithAssignments()
	fmt.Println("----------------------------------: Assignmants :---------------------------------->")
}
