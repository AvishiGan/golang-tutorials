package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 28,
	}

	fmt.Printf("Before: %v\n", ages)

	ages["Charlie"] = 32

	fmt.Printf("After: %v\n", ages)
}
