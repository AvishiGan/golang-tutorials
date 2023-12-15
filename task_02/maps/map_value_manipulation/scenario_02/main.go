package main

import "fmt"

func main() {

	counters := make(map[string]int)

	counters["apple"] = 50
	counters["mango"] = 60
	counters["orange"] = 70

	fmt.Printf("Before: %v\n", counters)

	value, ok := counters["apple"]

	if ok {
		counters["apple"] = value + 1
	}

	fmt.Printf("After: %v\n", counters)
}
