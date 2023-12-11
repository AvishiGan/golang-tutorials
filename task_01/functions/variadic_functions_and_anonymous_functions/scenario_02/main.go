package main

import "fmt"

func main() {
	doubleNumber := func(n int) int {
		return n * 2
	}
	doubledNumber := doubleNumber(4)
	fmt.Printf("Doubled Number: %d\n", doubledNumber)
}
