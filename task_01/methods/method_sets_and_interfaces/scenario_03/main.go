package main

import "fmt"

type Number int

func isEven(n Number) {
	if n%2 == 0 {
		fmt.Printf("%d is an even number", n)
	} else {
		fmt.Printf("%d is an odd number", n)
	}
}

func main() {
	var num Number = 8574631
	isEven(num)
}
