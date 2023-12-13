package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := sumSlice(nums)
	fmt.Printf("%v\n", nums)
	fmt.Printf("Sum: %v", sum)
}

func sumSlice(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}
