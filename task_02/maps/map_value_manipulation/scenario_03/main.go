package main

import "fmt"

func main() {

	counters := make(map[string]int)

	counters["apple"] = 50
	counters["mango"] = 60
	counters["banana"] = 40

	fmt.Printf("Before: %v\n", counters)

	_, ok := counters["orange"]

	if !ok {
		counters["orange"] = 1
	}

	fmt.Printf("After: %v\n", counters)
}
