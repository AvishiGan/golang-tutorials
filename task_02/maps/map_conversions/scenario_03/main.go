package main

import "fmt"

func main() {
	days := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
	}

	// Creating a new slice
	var keys []int

	for key := range days {
		keys = append(keys, key)
	}
	fmt.Println(keys)
}
