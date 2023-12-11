package main

import (
	"fmt"
)

func main() {
	printUserInput()
}

func printUserInput() {
	fmt.Println("Please enter your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hi %s! I'm Go", name)
}
