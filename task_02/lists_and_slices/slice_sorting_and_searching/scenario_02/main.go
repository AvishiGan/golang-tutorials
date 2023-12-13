package main

import "fmt"

func main() {
	nums := []int{21, 36, 97, 15, 104, 86}
	target := 80
	fmt.Printf("Slice: %v\n", nums)
	search(nums, target)
}

func search(nums []int, target int) {
	for i, num := range nums {
		if num == target {
			fmt.Printf("Found %d at index %d\n", target, i)
		}
	}
	fmt.Printf("%d is not found\n", target)
}
