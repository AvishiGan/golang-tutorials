package main

import "fmt"

func main() {
	isEven(2)
}

func isEven(number int) bool {
	if number%2 == 0 {
		fmt.Printf("%d is an even number\n", number)
		return true
	} else {
		fmt.Printf("%d is an odd number\n", number)
		return false
	}
}
