package main

import (
	"fmt"
	"os"
)

func main() {
	// Scenario 01
	fmt.Println("Hello World!")

	// Scenario 02
	print()

	// Scenario 03
	printUserInput()

	// // Scenario 04
	// printCmd()
}

// Scenario 02
func print() {
	fmt.Println("Hello World!")
}

// Scenario 03
func printUserInput() {
	fmt.Println("Please enter your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hi %s! I'm Go", name)
}

// Scenario 04
func printCmd() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
