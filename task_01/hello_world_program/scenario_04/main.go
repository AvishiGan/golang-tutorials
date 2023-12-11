package main

import (
	"fmt"
	"os"
)

func main() {

	// Retrieves the cmd arguments that were passed when starting the program
	args := os.Args

	if len(args) == 1 {
		fmt.Println("Hello World!")
	} else {
		fmt.Println("Hello" + args[1] + "!") // Hello [argument]!
	}
}
