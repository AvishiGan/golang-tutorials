package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Slice: %v\n", nums)

	reverseSlice(nums)
	fmt.Printf("Reversed slice: %v\n", nums)
}

func reverseSlice(nums []int) []int {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
	return nums
}
