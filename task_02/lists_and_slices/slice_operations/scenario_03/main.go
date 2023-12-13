package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Slice before filtering: %v\n", nums)

	fmt.Printf("Slice after filtering even numbers: %v\n", filterEven(nums))
}

func filterEven(nums []int) []int {
	result := []int{}
	for _, num := range nums {
		if num%2 == 0 {
			result = append(result, num)
		}
	}
	return result
}
