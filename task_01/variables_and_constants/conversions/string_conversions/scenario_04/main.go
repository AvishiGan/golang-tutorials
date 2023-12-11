package main

import (
	"fmt"
)

func main() {
	isValid := true
	isValidString := fmt.Sprintf("%t", isValid)
	str_concat := "That is "
	fmt.Printf(str_concat + isValidString)
}
