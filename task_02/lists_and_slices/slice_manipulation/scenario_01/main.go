package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// Start index -> inclusive
	// End index -> exclusive
	subset := numbers[1:4]
	fmt.Println(subset)
}
