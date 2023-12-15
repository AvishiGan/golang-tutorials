package main

import (
	"fmt"
)

func main() {
	ages := make(map[string]int)

	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 28

	fmt.Printf("Before: %v\n", ages)

	fmt.Printf("After: %v\n", incrementMapValues(ages, 2))
}

func incrementMapValues(m map[string]int, increment int) map[string]int {

	for k, v := range m {

		m[k] = v + increment
	}

	return m
}
