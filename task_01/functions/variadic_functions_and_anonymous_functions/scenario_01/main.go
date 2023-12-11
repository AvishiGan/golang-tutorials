package main

import "fmt"

func main() {
	sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
}

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Printf("\nTotal = %d\n", total)
	return total
}
