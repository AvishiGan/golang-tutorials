package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num string = "1234"
	numInt, err := strconv.Atoi(num) // Atoi stands for ASCII to Integer
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The integer is %v\n", numInt)
	}
}
