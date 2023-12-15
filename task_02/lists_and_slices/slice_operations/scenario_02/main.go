package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	max := findMax(nums)
	fmt.Printf("%v\n", nums)
	fmt.Printf("Max: %v", max)
}

func findMax(numbers []int) int {

	max := numbers[0]

	for _, number := range numbers {
		if number > max {
			max = number

		}
	}
	return max
}
