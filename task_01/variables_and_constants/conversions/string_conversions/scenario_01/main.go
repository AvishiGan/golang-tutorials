package main

import (
	"fmt"
	"strconv"
)

func main() {
	var count int = 16
	var countString string = strconv.Itoa(count) // Itoa -> Converts an integer to a string in base 10
	var str_var string = "cats"
	fmt.Printf(countString + " " + str_var)
}
