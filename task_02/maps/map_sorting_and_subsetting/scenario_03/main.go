package main

import (
	"fmt"
	"sort"
)

type keyValue struct {
	key   string
	value int
}

func sortMapByValue(m map[string]int) []keyValue {
	var sorted []keyValue
	for k, v := range m {
		sorted = append(sorted, keyValue{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].value < sorted[j].value
	})
	return sorted
}

func main() {
	ages := make(map[string]int)

	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 28

	fmt.Printf("Before sorting: %v\n", ages)

	sortedAges := sortMapByValue(ages)

	fmt.Printf("After sorting: %v\n", sortedAges)
}
