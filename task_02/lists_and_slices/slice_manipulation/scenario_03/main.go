package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("Numbers slice before: %v\n", numbers)

	newNumbers := []int{8, 9, 10}
	copy(numbers[len(numbers)-3:], newNumbers)
	fmt.Printf("Numbers slice after: %v\n", numbers)
}
