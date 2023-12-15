package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{25, 36, 49, 85, 10, 36, 74}
	fmt.Printf("Slice before sorting: %v\n", nums)

	sort.Ints(nums)
	fmt.Printf("Slice after sorting: %v\n", nums)
}
