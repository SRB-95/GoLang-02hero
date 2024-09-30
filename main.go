package main

import (
	"fmt"
	"golang-02hero/commandLineArgs"
	"golang-02hero/dataTypes"
	"golang-02hero/functions"
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

	fmt.Println("function: PlayWithCommandLineArgs")
	commandLineArgs.PlayWithCommandLineArgs()
	fmt.Println("-------------------------------: Command Line Args :------------------------------->")

	fmt.Println("function: playWithAssignments")
	playWithAssignments()
	fmt.Println("----------------------------------: Assignmants :---------------------------------->")
}
