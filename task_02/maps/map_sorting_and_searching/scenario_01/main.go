package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"Bob":     30,
		"Alice":   25,
		"Charlie": 28,
	}

	// []string -> type of the slice, 0 -> length, len(ages) -> capacity
	keys := make([]string, 0, len(ages))

	for key := range ages {
		keys = append(keys, key)
	}

	// Sort the keys alphabetically
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%s: %d\n", key, ages[key])
	}
}
